package services

import (
	"context"
	"encoding/json"
	bookingModels "github.com/Ddarli/gym/bookingservice/models"
	"github.com/Ddarli/gym/common/logger"
	"github.com/Ddarli/gym/kafka"
	"github.com/Ddarli/gym/shceduleservice/repository"
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
	"strconv"
)

type KafkaService struct {
	db     repository.ScheduleRepository
	logger *zap.SugaredLogger
}

var brokers = []string{"localhost:9095", "localhost:9096", "localhost:9097"}

func NewKafkaService(db repository.ScheduleRepository) *KafkaService {
	return &KafkaService{db: db, logger: logger.GetLogger()}
}

func (ks *KafkaService) checkScheduleAvailability(classId int) bool {
	_, err := ks.db.GetById(classId)
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
	id, _ := strconv.Atoi(msg.ScheduledClassId)
	hasSpace := ks.checkScheduleAvailability(id)
	if hasSpace {
		ks.logger.Infof("Schedule is correct")
		msg.Status = "3"
		producer, _ := kafka.NewSyncProducer(brokers)
		event, _ := json.Marshal(msg)
		err = kafka.SendMessage(ctx, producer, "schedule-checked-events", event)
		if err != nil {
			ks.logger.Warnf("Failed to send message to Kafka topic: %v", err)
			return err
		}
		return nil
	}
	return err
}
