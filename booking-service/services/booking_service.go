package services

import (
	"context"
	"github.com/Ddarli/gym/bookingservice/models"
	"github.com/Ddarli/gym/bookingservice/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Service struct {
	models.UnimplementedBookingServiceServer
	repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) models.BookingServiceServer {
	return &Service{repo: repo}
}

func (s *Service) CreateBooking(ctx context.Context, in *models.CreateBookingRequest) (*models.Booking, error) {
	booking := models.Booking{
		UserId:           in.UserId,
		ScheduledClassId: in.ScheduledClassId,
		BookingTime:      timestamppb.New(time.Now()),
		Status:           1,
	}

	err := s.repo.Create(&booking)
	if err != nil {
		return nil, err
	}
	return &booking, nil
}
func (s *Service) GetBooking(ctx context.Context, in *models.GetBookingRequest) (*models.Booking, error) {
	booking, err := s.repo.Get(in.GetId())
	if err != nil {
		return nil, err
	}
	return booking, nil
}
func (s *Service) CancelBooking(ctx context.Context, in *models.CancelBookingRequest) (*models.CancelBookingResponse, error) {
	res := models.CancelBookingResponse{}
	err := s.repo.Delete(in.GetId())
	if err != nil {
		res.Success = false
	}
	res.Success = true
	return &res, err
}
