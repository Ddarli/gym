package kafka

import "github.com/IBM/sarama"

func NewSyncProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	return sarama.NewSyncProducer(brokers, config)
}

func SendMessage(producer sarama.SyncProducer, topic string, message []byte) error {
	msg := sarama.ProducerMessage{Topic: topic, Value: sarama.ByteEncoder(message), Partition: -1}
	_, _, err := producer.SendMessage(&msg)
	return err
}
