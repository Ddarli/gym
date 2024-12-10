package services

import (
	"context"
	"encoding/json"
	bookingModels "github.com/Ddarli/gym/bookingservice/models"
	"github.com/Ddarli/gym/classservice/repository"
	"github.com/Ddarli/gym/common/logger"
	"github.com/Ddarli/gym/kafka"
	"github.com/Shopify/sarama"
	"go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"strconv"
)

type KafkaService struct {
	db     repository.ClassRepository
	logger *zap.SugaredLogger
}

var brokers = []string{"localhost:9095", "localhost:9096", "localhost:9097"}

func NewKafkaService(db repository.ClassRepository) *KafkaService {
	return &KafkaService{db: db, logger: logger.GetLogger()}
}

func (ks *KafkaService) checkClassAvailability(classId int) bool {
	class, err := ks.db.Get(classId)
	if err != nil {
		return false
	}
	if class.Capacity > 0 {
		return true
	}
	return false
}

func (ks *KafkaService) ProcessAvailabilityCheck(ctx context.Context, message *sarama.ConsumerMessage) error {
	var msg bookingModels.Booking
	err := json.Unmarshal(message.Value, &msg)
	if err != nil {
		return err
	}
	id, _ := strconv.Atoi(msg.ScheduledClassId)
	hasSpace := ks.checkClassAvailability(id)
	if hasSpace {
		ks.logger.Infof("Class has available spots")
		msg.Status = "2"
		producer, _ := kafka.NewSyncProducer(brokers)
		event, _ := json.Marshal(msg)

		kafkaMessage := &sarama.ProducerMessage{
			Topic: "class-checked-events",
			Value: sarama.ByteEncoder(event),
		}

		otel.GetTextMapPropagator().Inject(ctx, otelsarama.NewProducerMessageCarrier(kafkaMessage))

		_, _, err = producer.SendMessage(kafkaMessage)
		if err != nil {
			ks.logger.Warnf("Failed to send message to Kafka topic: %v", err)
			return err
		}
		return nil
	}
	return err
}
