package services

import (
	"context"
	"github.com/Ddarli/gym/bookingservice/models"
	"github.com/Ddarli/gym/bookingservice/repository"
	"strconv"
	"time"
)

type Service struct {
	models.UnimplementedBookingServiceServer
	repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) models.BookingServiceServer {
	return &Service{repo: repo}
}

func (s *Service) CreateBooking(ctx context.Context, in *models.CreateBookingRequest) (*models.CreateBookingResponse, error) {
	userId, _ := strconv.Atoi(in.UserId)
	classId, _ := strconv.Atoi(in.ScheduledClassId)
	booking := models.BookingModel{
		UserId:           userId,
		ScheduledClassId: classId,
		BookingTime:      time.Now(),
		Status:           1,
	}

	err := s.repo.Create(&booking)
	if err != nil {
		return nil, err
	}
	response := models.CreateBookingResponse{
		Booking: models.ToProto(&booking),
	}
	return &response, nil
}
func (s *Service) GetBooking(ctx context.Context, in *models.GetBookingRequest) (*models.GetBookingResponse, error) {
	id, _ := strconv.Atoi(in.GetId())
	booking, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	response := models.GetBookingResponse{
		Booking: models.ToProto(booking),
	}
	return &response, nil
}

func (s *Service) UpdateBooking(ctx context.Context, in *models.UpdateBookingRequest) (*models.UpdateBookingResponse, error) {
	return nil, nil
}

func (s *Service) DeleteBooking(ctx context.Context, in *models.DeleteBookingRequest) (*models.DeleteBookingResponse, error) {
	id, _ := strconv.Atoi(in.GetId())
	response := models.DeleteBookingResponse{
		Success: true,
	}
	err := s.repo.Delete(id)
	if err != nil {
		response.Success = false
		return &response, err
	}
	return &response, nil
}
