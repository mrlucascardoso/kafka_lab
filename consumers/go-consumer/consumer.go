package main

import "os"
import "fmt"
import "context"
import "time"
import kafka "github.com/segmentio/kafka-go"

func main() {
    // Wait for kafka to be ready
    time.Sleep(20 * time.Second)

    // Makes a new reader that consumes from our topic, partition 0, at offset 0.
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers:   []string{"kafka:29092"},
        Topic:     os.Getenv("KAFKA_TOPIC"),
        Partition: 0,
        MinBytes:  10e3, // 10KB
        MaxBytes:  10e6, // 10MB
    })
    r.SetOffset(0)

    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            break
        }
        fmt.Printf("Message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
    }

    r.Close()
}