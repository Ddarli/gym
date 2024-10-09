package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama"
	"go.uber.org/zap"
	"log"
)

type consumeFunction func(ctx context.Context, message *sarama.ConsumerMessage) error

type consumerGroupHandler struct {
	fn consumeFunction
}

func (c *consumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error   { return nil }
func (c *consumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error { return nil }
func (c *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		ctx := context.TODO()
		err := c.fn(ctx, message)
		if err != nil {
			log.Printf("Error handling message: %s\n", err)
		} else {
			session.MarkMessage(message, "")
		}
	}
	return nil
}
func StartConsuming(ctx context.Context, brokers []string, topic string, group string, consumeFunction consumeFunction,
	logger *zap.SugaredLogger) error {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		return err
	}
	h := otelsarama.WrapConsumerGroupHandler(&consumerGroupHandler{fn: consumeFunction})
	logger.Infof("Starting consumerGroupHandler group %s\n", group)
	go func() {
		for {
			if err := consumerGroup.Consume(ctx, []string{topic}, h); err != nil {
				fmt.Printf("Error from consumerGroupHandler: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	return nil
}
