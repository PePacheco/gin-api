package services

import (
	"notes-app/models"
	"notes-app/repositories"
	"testing"
)

func TestCreateNoteService_Execute(t *testing.T) {
	notes := []models.Note{
		{ID: 1, Name: "Note 1", Content: "Content for note 1"},
		{ID: 2, Name: "Note 2", Content: "Content for note 2"},
	}

	repo := &repositories.FakeNotesRepository{
		Notes: notes,
	}
	service := &CreateNoteService{Repository: repo}

	err := service.Execute(&models.Note{Name: "Note 3", Content: "Content for note 3"})

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if len(repo.Notes) != 3 {
		t.Errorf("Expected 3, got %v", len(repo.Notes))
	}
}
