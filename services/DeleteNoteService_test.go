package services

import (
	"notes-app/models"
	"notes-app/repositories"
	"testing"
)

func TestDeleteNoteService_Execute(t *testing.T) {
	notes := []models.Note{
		{ID: 1, Name: "Note 1", Content: "Content for note 1"},
		{ID: 2, Name: "Note 2", Content: "Content for note 2"},
	}

	repo := &repositories.FakeNotesRepository{
		Notes: notes,
	}
	service := &DeleteNoteService{Repository: repo}

	returnedNotes := service.Execute(notes[0].ID)

	if len(repo.Notes) != 1 {
		t.Fatalf("Expected notes %+v, but got %+v", notes, returnedNotes)
	}
}
