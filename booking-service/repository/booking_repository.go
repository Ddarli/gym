package repository

import "github.com/Ddarli/gym/bookingservice/models"

type BookingRepository interface {
	Create(booking *models.Booking) error
	Get(id string) (*models.Booking, error)
	Delete(id string) error
}
