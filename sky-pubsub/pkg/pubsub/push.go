package pubsub

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/pubsub"
)

func PushHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse the Pub/Sub message from the request body.
	message, err := pubsub.DecodeMessage(r.Body)
	if err != nil {
		http.Error(w, "failed to decode message: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the message.
	fmt.Printf("Received message: %s\n", string(message.Data))

	// Acknowledge the message so it isn't delivered again.
	message.Ack()

	// You can also call message.Nack() to signal that the message could not be processed
	// and should be retried later.
}

func CreatePushSubscription(ctx context.Context, projectID, topicName, subscriptionName, endpoint string) error {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("failed to create client: %v", err)
	}

	topic := client.Topic(topicName)
	exists, err := topic.Exists(ctx)
	if err != nil {
		return fmt.Errorf("failed to check if topic exists: %v", err)
	}

	if !exists {
		_, err = client.CreateTopic(ctx, topicName)
		if err != nil {
			return fmt.Errorf("failed to create topic: %v", err)
		}
	}

	// Create the subscription with a push endpoint.
	sub, err := client.CreateSubscription(ctx, subscriptionName, pubsub.SubscriptionConfig{
		Topic:      topic,
		PushConfig: pubsub.PushConfig{Endpoint: endpoint},
	})
	if err != nil {
		return fmt.Errorf("failed to create subscription: %v", err)
	}

	fmt.Printf("Created subscription %s with push endpoint %s\n", sub.ID(), endpoint)

	return nil
}
