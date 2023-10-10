package services

import (
	"fmt"
	"notes-app/models"
	"notes-app/repositories"
)

type UpdateNoteService struct {
	Repository repositories.NotesRepository
}

func (self *UpdateNoteService) Execute(id uint, input *models.Note) error {
	existingNote := self.Repository.ShowByName(input.Name)
	if existingNote != nil {
		return fmt.Errorf("Note with name %s already exists", input.Name)
	}
	self.Repository.Update(id, input)
	return nil
}
