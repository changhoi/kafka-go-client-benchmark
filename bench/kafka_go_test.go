package bench_test

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"runtime"
	"testing"
)

func messageBuilder(n int) kafka.Message {
	return kafka.Message{
		Value: []byte(fmt.Sprintf("%d-message", n)),
	}
}

func BenchmarkKafkaGo(b *testing.B) {
	w := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "benchtest",
		Async: true,
	}

	b.SetParallelism(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := w.WriteMessages(context.Background(), messageBuilder(i)); err != nil {
			log.Fatal(err)
		}
	}
	if err := w.Close(); err != nil {
		fmt.Println(err)
	}
}

func BenchmarkKafkaGoAsync(b *testing.B) {
	w := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "benchtest",
		Async: true,
	}

	b.SetParallelism(runtime.NumCPU())
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := w.WriteMessages(context.Background(), messageBuilder(0)); err != nil {
				log.Println(err.Error())
			}
		}
	})

	if err := w.Close(); err != nil {
		fmt.Println(err)
	}
}
