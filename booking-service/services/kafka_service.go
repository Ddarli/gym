package services

import (
	"context"
	"encoding/json"
	bookingModels "github.com/Ddarli/gym/bookingservice/models"
	"github.com/Ddarli/gym/bookingservice/repository"
	"github.com/Shopify/sarama"
	"go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama"
	"go.opentelemetry.io/otel"
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

func (ks *KafkaService) updateBooking(ctx context.Context, bookingId int) bool {
	booking, _ := ks.db.Get(ctx, bookingId)
	booking.Status = 3

	err := ks.db.Update(ctx, bookingId, booking)
	if err != nil {
		return false
	}

	return true
}

func (ks *KafkaService) ProcessAvailabilityCheck(_ context.Context, message *sarama.ConsumerMessage) error {
	var msg bookingModels.Booking

	carrier := otelsarama.NewConsumerMessageCarrier(message)
	ctx := otel.GetTextMapPropagator().Extract(context.Background(), carrier)

	tracer := otel.Tracer("kafka-consumer")
	ctx, span := tracer.Start(ctx, "consume-kafka-message")

	err := json.Unmarshal(message.Value, &msg)
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(msg.GetId())
	res := ks.updateBooking(ctx, id)
	if res {
		return nil
	}

	span.End()

	return err
}
