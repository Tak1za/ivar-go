package client

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func GetFirestoreClient() (*firestore.Client, error) {
	// opt := option.WithCredentialsFile("../ivar-cred.json")
	opt := option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_JSON")))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("Error initializing the firebase client: %s", err)
		return nil, err
	}

	fc, _ := app.Firestore(context.Background())
	return fc, nil
}

//func VerifyAccessToken(idToken string) (bool, error) {
//	// opt := option.WithCredentialsFile("../ivar-cred.json")
//	opt := option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_JSON")))
//	app, err := firebase.NewApp(context.Background(), nil, opt)
//	if err != nil {
//		log.Printf("Error initializing the firebase client: %s", err)
//		return false, err
//	}
//
//	client, err := app.Auth(context.Background())
//	if err != nil {
//		log.Printf("error getting Auth client: %v\n", err)
//		return false, err
//	}
//
//	_, err = client.VerifyIDToken(context.Background(), idToken)
//	if err != nil {
//		log.Printf("error verifying ID token: %v\n", err)
//		return false, err
//	}
//
//	return true, nil
//}
