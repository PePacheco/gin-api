package services

import (
	"notes-app/models"
	"notes-app/repositories"
)

type ListAllNotesService struct {
	Repository repositories.NotesRepository
}

func (s *ListAllNotesService) Execute() []models.Note {
	return s.Repository.All()
}
