package repository

import "github.com/Ddarli/gym/userservice/models"

type UserRepository interface {
	Create(user *models.User) error
	GetById(userId string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
}
