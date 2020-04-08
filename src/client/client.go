package client

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

func GetFirestoreClient() (*firestore.Client, error) {
	opt := option.WithCredentialsFile("src/ivar-cred.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing the firebase client: %s", err)
	}
	fc, _ := app.Firestore(context.Background())
	return fc, nil
}
