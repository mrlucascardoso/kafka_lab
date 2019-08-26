package main

import "os"
import "fmt"
import "context"
import "time"
import kafka "github.com/segmentio/kafka-go"

func main() {
    // Wait for kafka to be ready
    time.Sleep(20 * time.Second)

	if os.Getenv("KAFKA_API_LEVEL") == "Low" {
        // Makes a new reader that consumes from our topic, partition 0, at offset 0.
        r := kafka.NewReader(kafka.ReaderConfig{
            Brokers:   []string{"kafka:29092"},
            Topic:     os.Getenv("KAFKA_TOPIC"),
            Partition: 0, // Since we are defining a Low Level API we need to set the partition that we are connecting to.
            MinBytes:  10e3, // 10KB
            MaxBytes:  10e6, // 10MB
        })
        r.SetOffset(0) // On Low Level API we also need to set the offset.

        for {
            m, err := r.ReadMessage(context.Background())
            if err != nil {
                break
            }
            fmt.Printf("Go consumer received value using Low level API: %s\n", string(m.Value))
        }

        r.Close()
	}

	if os.Getenv("KAFKA_API_LEVEL") == "High" {
        // Makes a new reader that consumes from our topic.
        r := kafka.NewReader(kafka.ReaderConfig{
            Brokers:   []string{"kafka:29092"},
            GroupID:   "goConsumerGroup",
            Topic:     os.Getenv("KAFKA_TOPIC"),
            MinBytes:  10e3, // 10KB
            MaxBytes:  10e6, // 10MB
        })

        for {
            m, err := r.ReadMessage(context.Background())
            if err != nil {
                break
            }
            fmt.Printf("Go consumer received value using High level API: %s\n", string(m.Value))
        }

        r.Close()
	}
}