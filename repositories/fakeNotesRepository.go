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

func (f *FakeNotesRepository) DeleteById(id uint) {
	for i, note := range f.Notes {
		if note.ID == id {
			f.Notes = append(f.Notes[:i], f.Notes[i+1:]...)
		}
	}
}

func (f *FakeNotesRepository) Update(id uint, note *models.Note) {
	for i, n := range f.Notes {
		if n.ID == id {
			f.Notes[i] = *note
		}
	}
}
