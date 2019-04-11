
package main

import (
    "github.com/segmentio/kafka-go"
    "context"
)

func main() {

// make a writer that produces to topic-A, using the least-bytes distribution
w := kafka.NewWriter(kafka.WriterConfig{
	Brokers: []string{"172.17.0.1:9092"},
	Topic:   "error",
	Balancer: &kafka.LeastBytes{},
})

w.WriteMessages(context.Background(),
	kafka.Message{
		Key:   []byte("Key-A"),
		Value: []byte("Hello World!"),
	},
	kafka.Message{
		Key:   []byte("Key-B"),
		Value: []byte("One!"),
	},
	kafka.Message{
		Key:   []byte("Key-C"),
		Value: []byte("Two!"),
	},
)

w.Close()

}