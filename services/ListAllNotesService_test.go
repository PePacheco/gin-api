package services

import (
	"notes-app/models"
	"notes-app/repositories"
	"reflect"
	"testing"
)

func TestListAllNotesService_Execute(t *testing.T) {
	notes := []models.Note{
		{ID: 1, Name: "Note 1", Content: "Content for note 1"},
		{ID: 2, Name: "Note 2", Content: "Content for note 2"},
	}

	repo := &repositories.FakeNotesRepository{
		Notes: notes,
	}
	service := &ListAllNotesService{Repository: repo}

	returnedNotes := service.Execute()

	if !reflect.DeepEqual(returnedNotes, notes) {
		t.Fatalf("Expected notes %+v, but got %+v", notes, returnedNotes)
	}
}
