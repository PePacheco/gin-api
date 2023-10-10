package services

import (
	"notes-app/models"
	"notes-app/repositories"
)

type ListAllNotesService struct {
	Repository repositories.NotesRepository
}

func (self *ListAllNotesService) Execute() []models.Note {
	return self.Repository.All()
}
