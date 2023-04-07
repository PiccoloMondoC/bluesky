package config

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub"
)

// Define the necessary settings for Pub/Sub topics and subscriptions.
const (
	topicName                  = "example-topic"
	messageRetentionDuration   = 24 * time.Hour
	messageOrdering            = true
	messageSizeLimit           = 10 * 1024 * 1024  // 10MB
	messageDeliveryOptions     = pubsub.DeliverAll // or pubsub.DeliverOnce
	messageAcknowledgement     = true
	messageFiltering           = true
	topicAccessControl         = pubsub.Publisher
	subscriptionAccessControl  = pubsub.Subscriber
	subscriptionReceiveTimeout = 10 * time.Second
)

// Create a Pub/Sub topic.
func createTopic(ctx context.Context, client *pubsub.Client) (*pubsub.Topic, error) {
	topic := client.Topic(topicName)
	exists, err := topic.Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		_, err := client.CreateTopic(ctx, topicName)
		if err != nil {
			return nil, err
		}
	}
	topic, err = client.Topic(topicName).Update(ctx, pubsub.TopicConfig{
		MessageRetentionDuration: messageRetentionDuration,
		OrderingKey:              messageOrdering,
	})
	if err != nil {
		return nil, err
	}
	return topic, nil
}

// Create a Pub/Sub subscription.
func createSubscription(ctx context.Context, client *pubsub.Client, topic *pubsub.Topic, subscriptionName string) (*pubsub.Subscription, error) {
	subscription := client.Subscription(subscriptionName)
	exists, err := subscription.Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		_, err := client.CreateSubscription(ctx, subscriptionName, pubsub.SubscriptionConfig{
			Topic:       topic,
			AckDeadline: subscriptionReceiveTimeout,
		})
		if err != nil {
			return nil, err
		}
	}
	subscription, err = client.Subscription(subscriptionName).Update(ctx, pubsub.SubscriptionConfig{
		AckDeadline: subscriptionReceiveTimeout,
	})
	if err != nil {
		return nil, err
	}
	return subscription, nil
}
