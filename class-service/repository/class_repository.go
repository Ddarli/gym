package repository

import "github.com/Ddarli/gym/classservice/models"

type ClassRepository interface {
	Create(*models.ClassModel) (*models.ClassModel, error)
	Get(int) (*models.ClassModel, error)
	GetAll() ([]*models.ClassModel, error)
}
