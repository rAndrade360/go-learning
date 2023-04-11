package main

import (
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "gogroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatal("Err to start consumer: ", err.Error())
	}

	err = consumer.Subscribe("kafka-labs", nil)
	if err != nil {
		log.Fatal("Err to start consumer: ", err.Error())
	}

	for {
		m, err := consumer.ReadMessage(time.Second * 2)
		if err != nil {
			log.Println("Err to read message: ", err.Error())
			continue
		}

		log.Println(m.String())
		log.Println(string(m.Value))
		log.Println(m.TopicPartition.Partition)

		log.Println(consumer.String())
	}
}
