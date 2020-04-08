package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/iterator"
	"ivar-go/src/client"
	"ivar-go/src/models"
	"log"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Welcome to the home page!")
}

func UsersController(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	firestore, _ := client.GetFirestoreClient()
	defer firestore.Close()

	iter := firestore.Collection("users").Documents(ctx)
	var users []models.User
	for {
		doc, err := iter.Next()
		var user models.User
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		jsonString, _ := json.Marshal(doc.Data())
		err = json.Unmarshal(jsonString, &user)
		if err != nil {
			log.Fatalf("Error unmarshalling to json: %s", err)
		}
		user.ID = doc.Ref.ID
		users = append(users, user)
	}

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatalf("Error encoding data: %s", err)
	}
}
