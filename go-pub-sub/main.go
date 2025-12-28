package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Publisher
	go func() {
		time.Sleep(2 * time.Second)
		err := client.Publish(ctx, "my-channel", "Hello from Go!").Err()
		if err != nil {
			log.Fatalf("Error publishing message: %v", err)
		}
		fmt.Println("Message published!")
	}()

	// Subscriber
	sub := client.Subscribe(ctx, "my-channel")
	defer sub.Close()

	// Wait for confirmation that subscription is created
	_, err := sub.Receive(ctx)
	if err != nil {
		log.Fatalf("Error subscribing: %v", err)
	}

	// Listen for messages in the background
	ch := sub.Channel()

	// Handle received messages
	for msg := range ch {
		fmt.Printf("Received message: %s\n", msg.Payload)
	}
}
