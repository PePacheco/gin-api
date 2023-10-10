package services

import (
	"fmt"
	"notes-app/repositories"
)

type DeleteNoteService struct {
	Repository repositories.NotesRepository
}

func (self *DeleteNoteService) Execute(id uint) error {
	existingNote := self.Repository.ShowById(id)
	if existingNote == nil {
		return fmt.Errorf("Note with id %d not found", id)
	}
	self.Repository.DeleteById(id)
	return nil
}
