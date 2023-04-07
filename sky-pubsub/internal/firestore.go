package data

import (
	"context"

	"firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

func CreateFirestoreClient(ctx context.Context, projectID string) (*db.Client, error) {
	cred := option.WithCredentialsFile("/path/to/service-account-key.json")
	app, err := firebase.NewApp(ctx, nil, cred)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
