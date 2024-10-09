package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama"
	"go.opentelemetry.io/otel"
)

func NewSyncProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return otelsarama.WrapSyncProducer(config, producer), nil
}

func SendMessage(ctx context.Context, producer sarama.SyncProducer, topic string, message []byte) error {
	tr := otel.Tracer("kafka-producer")
	ctx, span := tr.Start(ctx, "SendMessage")
	defer span.End()
	msg := &sarama.ProducerMessage{Topic: topic, Value: sarama.ByteEncoder(message), Partition: -1}
	otel.GetTextMapPropagator().Inject(ctx, otelsarama.NewProducerMessageCarrier(msg))
	_, _, err := producer.SendMessage(msg)
	if err != nil {
		span.RecordError(err)
	}
	return err
}
