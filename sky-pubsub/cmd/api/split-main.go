package main

import (
	"log"
)

func main() {
	// Load the PubSub Manager configuration from environment variables
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create a new PubSub Manager instance
	mgr, err := NewPubSubManager(cfg)
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub Manager: %v", err)
	}

	// Start the Pub/Sub Manager instance
	if err := mgr.Start(); err != nil {
		log.Fatalf("Failed to start Pub/Sub Manager: %v", err)
	}
}
