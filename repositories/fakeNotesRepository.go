package repositories

import "notes-app/models"

type FakeNotesRepository struct {
	Notes []models.Note
}

func (f *FakeNotesRepository) All() []models.Note {
	return f.Notes
}

func (f *FakeNotesRepository) Create(note *models.Note) {
	f.Notes = append(f.Notes, *note)
}
