package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

var topic = "benchtest"

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":     "127.0.0.1:9092",
		"security.protocol":     "PLAINTEXT",
		"broker.address.family": "v4",
	})

	if err != nil {
		panic(err)
	}

	ch := p.ProduceChannel()

	go func() {
		for v := range p.Events() {
			fmt.Printf("%+v", v)
		}
	}()

	for {
		time.Sleep(time.Second)
		ch <- &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: nil,
		}
	}
}
