package main

import (
	"fmt"
	"log"

	"github.com/KaueSabinoSRV17/DeliveryGoSimulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {

	messageChannel := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(messageChannel)
	go consumer.Consume()

	for msg := range messageChannel {
		fmt.Println(string(msg.Value))
	}

	// rt := route.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }
	// rt.LoadPositions()

	// stringJson, _ := rt.ExportJsonPositions()
	// fmt.Println(stringJson[1])
}
