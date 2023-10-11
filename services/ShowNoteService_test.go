package services

import (
	"notes-app/models"
	"notes-app/repositories"
	"reflect"
	"testing"
)

func TestShowNoteService_Execute(t *testing.T) {
	notes := []models.Note{
		{ID: 1, Name: "Note 1", Content: "Content for note 1"},
		{ID: 2, Name: "Note 2", Content: "Content for note 2"},
	}

	repo := &repositories.FakeNotesRepository{
		Notes: notes,
	}
	service := &ShowNoteService{Repository: repo}

	returnedNote := service.Execute(notes[0].ID)

	if !reflect.DeepEqual(*returnedNote, notes[0]) {
		t.Fatalf("Expected notes %+v, but got %+v", notes, *returnedNote)
	}
}
