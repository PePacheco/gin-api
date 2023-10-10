package services

import (
	"fmt"
	"notes-app/models"
	"notes-app/repositories"
)

type CreateNoteService struct {
	Repository repositories.NotesRepository
}

func (self *CreateNoteService) Execute(input *models.Note) error {
	existingNote := self.Repository.ShowByName(input.Name)
	if existingNote != nil {
		return fmt.Errorf("Note with name %s already exists", input.Name)
	}
	self.Repository.Create(input)
	return nil
}
