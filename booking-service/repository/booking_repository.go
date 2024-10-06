package repository

import "github.com/Ddarli/gym/bookingservice/models"

type BookingRepository interface {
	Create(*models.BookingModel) error
	Get(int) (*models.BookingModel, error)
	Update(int, *models.BookingModel) error
	Delete(int) error
}
