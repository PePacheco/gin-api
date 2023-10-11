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
