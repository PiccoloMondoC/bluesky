package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
)

func ReceiveMessages(ctx context.Context, projectID, subID string) error {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("failed to create client: %v", err)
	}

	sub := client.Subscription(subID)
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Got message: %q\n", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("failed to receive messages: %v", err)
	}

	return nil
}
