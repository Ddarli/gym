package repository

import "github.com/Ddarli/gym/trainerservice/models"

type TrainerRepository interface {
	Create(*models.TrainerModel) (*models.TrainerModel, error)
	Get(int) (*models.TrainerModel, error)
}
