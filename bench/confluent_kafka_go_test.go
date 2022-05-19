package bench_test

import (
	"fmt"
	kafkac "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"runtime"
	"testing"
)

var (
	broker = "localhost:9092"
	topic  = "benchtest"
)

var (
	kafacConfigMap = kafkac.ConfigMap{
		"bootstrap.servers":     broker,
		"security.protocol":     "PLAINTEXT",
		"broker.address.family": "v4",
	}
)

func kafkacMessageBuilder(n int) *kafkac.Message {
	return &kafkac.Message{
		TopicPartition: kafkac.TopicPartition{
			Topic:     &topic,
			Partition: kafkac.PartitionAny,
		},
		Value: []byte(fmt.Sprintf("%d-mssage", n)),
	}
}

func BenchmarkConfluentKafkaGoAsync(b *testing.B) {
	b.ReportAllocs()
	p, err := kafkac.NewProducer(&kafacConfigMap)

	if err != nil {
		log.Fatal(err)
	}

	ch := p.ProduceChannel()

	go func() {
		for range p.Events() {
		}
	}()

	b.SetParallelism(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch <- kafkacMessageBuilder(i)
	}
	p.Flush(1000 * 60)
	b.StopTimer()
}

func BenchmarkConfluentKafkaGoAsyncInParalell(b *testing.B) {
	b.ReportAllocs()
	p, err := kafkac.NewProducer(&kafacConfigMap)
	if err != nil {
		log.Fatal(err)
	}

	ch := p.ProduceChannel()

	go func() {
		for range p.Events() {
		}
	}()

	b.ResetTimer()
	b.SetParallelism(runtime.NumCPU())
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ch <- kafkacMessageBuilder(0)
		}
	})
	p.Flush(1000 * 60)
	b.StopTimer()
}
