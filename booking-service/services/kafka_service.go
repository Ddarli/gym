package services

import (
	"context"
	"encoding/json"
	bookingModels "github.com/Ddarli/gym/bookingservice/models"
	"github.com/Ddarli/gym/bookingservice/repository"
	"github.com/Ddarli/gym/common/logger"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"strconv"
)

type KafkaService struct {
	db     repository.BookingRepository
	logger *zap.SugaredLogger
}

func NewKafkaService(db repository.BookingRepository) *KafkaService {
	return &KafkaService{db: db, logger: logger.GetLogger()}
}

func (ks *KafkaService) updateBooking(bookingId int) bool {
	booking, _ := ks.db.Get(bookingId)
	booking.Status = 3
	err := ks.db.Update(bookingId, booking)
	if err != nil {
		return false
	}
	return true
}

func (ks *KafkaService) ProcessAvailabilityCheck(ctx context.Context, message *sarama.ConsumerMessage) error {
	var msg bookingModels.Booking
	err := json.Unmarshal(message.Value, &msg)
	if err != nil {
		return err
	}
	id, _ := strconv.Atoi(msg.GetId())
	res := ks.updateBooking(id)
	if res {
		return nil
	}
	return err
}
