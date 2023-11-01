package services

import (
	"notes-app/models"
	"notes-app/repositories"
	"testing"
)

func TestUpdateNoteService_Execute(t *testing.T) {
	notes := []models.Note{
		{ID: 1, Name: "Note 1", Content: "Content for note 1"},
		{ID: 2, Name: "Note 2", Content: "Content for note 2"},
	}

	repo := &repositories.FakeNotesRepository{
		Notes: notes,
	}
	service := &UpdateNoteService{Repository: repo}

	err := service.Execute(notes[0].ID, &models.Note{Name: "Note 1 Updated", Content: "Content for note 1 updated"})

	if err != nil {
		t.Fatalf("Expected no error, but got %+v", err)
	}
}

func TestUpdateNoteService_Execute_WithNameConflict(t *testing.T) {
	notes := []models.Note{
		{ID: 1, Name: "Note 1", Content: "Content for note 1"},
		{ID: 2, Name: "Note 2", Content: "Content for note 2"},
	}

	repo := &repositories.FakeNotesRepository{
		Notes: notes,
	}
	service := &UpdateNoteService{Repository: repo}

	expectedErrMsg := "Note with name Note 2 already exists"

	err := service.Execute(notes[0].ID, &models.Note{Name: "Note 2", Content: "Content for note 1 updated"})

	if err == nil {
		t.Fatalf("Expected an error due to name conflict, but got nil")
	}

	if err.Error() != expectedErrMsg {
		t.Fatalf("Expected error message to be %s, but got %s", expectedErrMsg, err.Error())
	}
}
