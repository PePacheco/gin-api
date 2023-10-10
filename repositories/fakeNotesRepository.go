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

func (f *FakeNotesRepository) ShowById(id uint) *models.Note {
	for _, note := range f.Notes {
		if note.ID == id {
			return &note
		}
	}
	return nil
}

func (f *FakeNotesRepository) ShowByName(name string) *models.Note {
	for _, note := range f.Notes {
		if note.Name == name {
			return &note
		}
	}
	return nil
}
