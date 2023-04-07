package pubsub

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

func FanOut(ctx context.Context, message *pubsub.Message) {
	// Get the topic name from the message attributes
	topic := message.Attributes["topic"]

	// Create a Pub/Sub client
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return
	}

	// Get a list of subscriptions associated with the topic
	subs := client.Topic(topic).Subscriptions(ctx)
	for {
		sub, err := subs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Failed to get subscription: %v", err)
			continue
		}

		// Publish the message to the subscription
		result := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
			// Handle the message
			log.Printf("Received message from subscription %s: %s", sub.ID(), string(msg.Data))
			msg.Ack()
		})
		<-result
	}
}
