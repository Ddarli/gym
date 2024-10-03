package repository

import "github.com/Ddarli/gym/userservice/models"

type UserRepository interface {
	Create(user *models.User) error
	Get(userId string) (*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
}
