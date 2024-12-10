package repository

import (
	"context"
	"github.com/Ddarli/gym/bookingservice/models"
)

type BookingRepository interface {
	Create(*models.BookingModel) (int, error)
	Get(context.Context, int) (*models.BookingModel, error)
	Update(context.Context, int, *models.BookingModel) error
	Delete(int) error
}
