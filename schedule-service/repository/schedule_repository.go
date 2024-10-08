package repository

import "github.com/Ddarli/gym/shceduleservice/models"

type ScheduleRepository interface {
	GetById(int) (models.ScheduleModel, error)
	Create(models.ScheduleModel) error
}
