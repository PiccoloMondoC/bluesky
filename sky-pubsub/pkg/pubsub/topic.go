package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func CreateTopic(ctx context.Context, projectID, topicName string) error {
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

	return nil
}
