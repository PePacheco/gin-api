package repositories

import "notes-app/models"

type NotesRepository interface {
	All() []models.Note
	Create(note *models.Note)
	ShowById(id uint) *models.Note
	ShowByName(name string, excludeId ...uint) *models.Note
	DeleteById(id uint)
	Update(id uint, note *models.Note)
}

type NotesRepositoryImpl struct{}

func (r *NotesRepositoryImpl) All() []models.Note {
	return *models.NotesAll()
}

func (r *NotesRepositoryImpl) Create(note *models.Note) {
	models.NotesCreate(note)
}

func (r *NotesRepositoryImpl) ShowById(id uint) *models.Note {
	return models.NoteById(id)
}

func (r *NotesRepositoryImpl) ShowByName(name string, excludeId ...uint) *models.Note {
	return models.NoteByName(name, excludeId...)
}

func (r *NotesRepositoryImpl) DeleteById(id uint) {
	note := models.NoteById(id)
	models.NotesDelete(note)
}

func (r *NotesRepositoryImpl) Update(id uint, note *models.Note) {
	models.NotesUpdate(id, note)
}
