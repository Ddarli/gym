package kafka

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

type consumeFunction func(ctx context.Context, message *sarama.ConsumerMessage) error

type consumer struct {
	fn consumeFunction
}

func (c *consumer) Setup(session sarama.ConsumerGroupSession) error   { return nil }
func (c *consumer) Cleanup(session sarama.ConsumerGroupSession) error { return nil }
func (c *consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
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
func StartConsuming(ctx context.Context, brokers []string, topic string, group string, consumeFunction consumeFunction) error {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		return err
	}

	consumer := consumer{
		fn: consumeFunction,
	}

	go func() {
		for {
			if err := consumerGroup.Consume(ctx, []string{topic}, &consumer); err != nil {
				fmt.Printf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	return nil
}
