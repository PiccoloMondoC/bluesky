package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/pubsub/v1"
	"google.golang.org/grpc"
)

// PubSubManager represents the Pub/Sub Manager instance.
type PubSubManager struct {
	Config     *Config
	PubSub     *pubsub.Service
	HttpServer *http.Server
}

// NewPubSubManager creates a new Pub/Sub Manager instance.
func NewPubSubManager(cfg *Config) (*PubSubManager, error) {
	ctx := context.Background()
	creds, err := grpc.GoogleDefaultCredentials()
	if err != nil {
		log.Fatalf("Failed to get gRPC credentials: %v", err)
	}

	// Authenticate with the service account key file
	opts := option.WithCredentialsFile(cfg.ServiceAccount)
	pubSub, err := pubsub.NewService(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Pub/Sub client: %v", err)
	}

	// Create the HTTP server and router
	router := mux.NewRouter()
	server := &http.Server{
		Addr:    cfg.PubsubEndpoint,
		Handler: router,
	}

	// Create the PubSub Manager instance
	return &PubSubManager{
		Config:     cfg,
		PubSub:     pubSub,
		HttpServer: server,
	}, nil
}

// Start starts the Pub/Sub Manager instance.
func (mgr *PubSubManager) Start() error {
	log.Printf("Starting Pub/Sub Manager on port %s", mgr.HttpServer.Addr)
	return mgr.HttpServer.ListenAndServe()
}

// Stop stops the Pub/Sub Manager instance.
func (mgr *PubSubManager) Stop() error {
	log.Printf("Stopping Pub/Sub Manager")
	return mgr.HttpServer.Shutdown(context.Background())
}
