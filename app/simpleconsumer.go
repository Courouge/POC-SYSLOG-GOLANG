
package main

import (
    "fmt"
    "github.com/segmentio/kafka-go"
    "context"
    "strings"
)

func main() {

// make a new reader that consumes from topic-A, partition 0, at offset 42
r := kafka.NewReader(kafka.ReaderConfig{
    Brokers:   []string{"localhost:9092"},
    Topic:     "events",
    Partition: 0,
    MinBytes:  10e3, // 10KB
    MaxBytes:  10e6, // 10MB
})


for {
    m, err := r.ReadMessage(context.Background())
    if err != nil {
        fmt.Printf("Kafka server is not connected")
        break
    }
    //fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
    stringSlice := strings.Split(string(m.Value), "|")
    fmt.Printf("%s\n",stringSlice)
}


r.Close()
}