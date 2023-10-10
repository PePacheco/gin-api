package services

import (
	"notes-app/models"
	"notes-app/repositories"
)

type CreateNoteService struct {
	Repository repositories.NotesRepository
}

func (s *CreateNoteService) Execute(input *models.Note) {
	s.Repository.Create(input)
}
