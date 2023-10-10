package repositories

import "notes-app/models"

type UsersRepository interface {
	Create(user *models.User) error
	ShowByUsername(username string) *models.User
}

type UsersRepositoryImpl struct{}

func (r *UsersRepositoryImpl) Create(user *models.User) error {
	return models.UserCreate(user)
}

func (r *UsersRepositoryImpl) ShowByUsername(username string) *models.User {
	return models.UserByUsername(username)
}
