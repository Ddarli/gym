package services

import (
	"context"
	"encoding/json"
	"github.com/Ddarli/gym/bookingservice/models"
	"github.com/Ddarli/gym/bookingservice/repository"
	"github.com/Ddarli/gym/kafka"
	"github.com/Shopify/sarama"
	"go.opentelemetry.io/otel"
	"log"
	"strconv"
	"time"
)

var brokers = []string{"localhost:9095", "localhost:9096", "localhost:9097"}

type Service struct {
	models.UnimplementedBookingServiceServer
	repo     repository.BookingRepository
	producer sarama.SyncProducer
}

func NewBookingService(repo repository.BookingRepository) models.BookingServiceServer {
	createdProducer, err := kafka.NewSyncProducer(brokers)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	return &Service{repo: repo, producer: createdProducer}
}

func (s *Service) sendBookingCreatedEvent(booking *models.Booking) error {
	messageBytes, err := json.Marshal(booking)
	if err != nil {
		return err
	}
	err = kafka.SendMessage(context.Background(), s.producer, "booking-event", messageBytes)
	if err != nil {
		return err
	}
	return nil
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

	id, err := s.repo.Create(&booking)
	if err != nil {
		return nil, err
	}
	booking.Id = id
	response := models.CreateBookingResponse{
		Booking: models.ToProto(&booking),
	}
	err = s.sendBookingCreatedEvent(models.ToProto(&booking))
	if err != nil {
		return nil, err
	}
	return &response, nil
}
func (s *Service) GetBooking(ctx context.Context, in *models.GetBookingRequest) (*models.GetBookingResponse, error) {
	tracer := otel.Tracer("booking-service")
	_, span := tracer.Start(ctx, "GetBooking")
	defer span.End()
	id, _ := strconv.Atoi(in.GetId())
	booking, err := s.repo.Get(ctx, id)
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
