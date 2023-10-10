package repositories

import "notes-app/models"

type NotesRepository interface {
	All() []models.Note
	Create(note *models.Note)
}

type NotesRepositoryImpl struct{}

func (r *NotesRepositoryImpl) All() []models.Note {
	return *models.NotesAll()
}

func (r *NotesRepositoryImpl) Create(note *models.Note) {
	models.NotesCreate(note)
}
