package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	MessageChanel chan *ckafka.Message
}

func (consumer *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}

	csmr, consumerError := ckafka.NewConsumer(configMap)
	if consumerError != nil {
		log.Fatal("Error consuming Kafka message:" + consumerError.Error())
	}

	topics := []string{os.Getenv("KafkaReadTopic")}
	csmr.SubscribeTopics(topics, nil)
	fmt.Println("Kafka consumer has been started")

	for {
		message, messageError := csmr.ReadMessage(-1)
		fmt.Println("Antes")
		if messageError == nil {
			fmt.Println("Entrou")
			consumer.MessageChanel <- message
		}
	}

}
