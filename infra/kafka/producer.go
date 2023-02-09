package kafka

import (
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaConsumer(messageChannel chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MessageChanel: messageChannel,
	}
}

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}

	producer, producerError := ckafka.NewProducer(configMap)
	if producerError != nil {
		log.Println(producerError.Error())
	}

	return producer

}

func Publish(msg string, topic string, producer *ckafka.Producer) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}
	messageError := producer.Produce(message, nil)
	if messageError != nil {
		return messageError
	}

	return nil
}
