package main

import (
    "context"
    "fmt"
    "log"

    "cloud.google.com/go/pubsub"
)

func main() {
    ctx := context.Background()

    // Set your project ID here
    projectID := "your-project-id"

    // Set the topic name here
    topicName := "your-topic-name"

    // Create a Pub/Sub client
    client, err := pubsub.NewClient(ctx, projectID)
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }

    // Create the topic
    topic, err := client.CreateTopic(ctx, topicName)
    if err != nil {
        log.Fatalf("Failed to create topic: %v", err)
    }

    // Print the topic name
    fmt.Printf("Topic created: %v\n", topicName)
}
