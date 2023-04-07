package main

import "os"

// LoadConfig loads the configuration settings from environment variables.
func LoadConfig() (*Config, error) {
	cfg := &Config{
		ProjectID:      os.Getenv("PUBSUB_MANAGER_PROJECT_ID"),
		PubsubTopic:    os.Getenv("PUBSUB_MANAGER_TOPIC"),
		PubsubSub:      os.Getenv("PUBSUB_MANAGER_SUBSCRIPTION"),
		PubsubEndpoint: os.Getenv("PUBSUB_MANAGER_ENDPOINT"),
	}

	return cfg, nil
}
