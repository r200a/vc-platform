package events

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

type ApplicationEvent struct {
	EventType     string `json:"event_type"`
	ApplicationID string `json:"application_id"`
	StartupID     string `json:"startup_id"`
	StartupName   string `json:"startup_name"`
	VCID          string `json:"vc_id"`
	VCName        string `json:"vc_name"`
	FounderEmail  string `json:"founder_email"`
	VCEmail       string `json:"vc_email"`
	Status        string `json:"status"`
	RejectionNote string `json:"rejection_note"`
}

type Producer struct {
	writer *kafka.Writer
}

func NewProducer() *Producer {
	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		broker = "kafka:9092"
	}

	writer := &kafka.Writer{
		Addr:         kafka.TCP(broker),
		Topic:        "application.events",
		Balancer:     &kafka.LeastBytes{},
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	return &Producer{writer: writer}
}

func (p *Producer) Publish(event ApplicationEvent) {
	data, err := json.Marshal(event)
	if err != nil {
		log.Printf("failed to marshal event: %v", err)
		return
	}

	err = p.writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(event.ApplicationID),
			Value: data,
		},
	)
	if err != nil {
		log.Printf("failed to publish event %s: %v", event.EventType, err)
		return
	}

	log.Printf("event published: %s for application %s",
		event.EventType, event.ApplicationID)
}

func (p *Producer) Close() {
	p.writer.Close()
}
