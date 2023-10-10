package services

import (
	"notes-app/models"
	"notes-app/repositories"
)

type ShowNoteService struct {
	Repository repositories.NotesRepository
}

func (self *ShowNoteService) Execute(id uint) *models.Note {
	return self.Repository.ShowById(id)
}
