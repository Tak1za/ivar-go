package controllers

import (
	"encoding/json"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"ivar-go/src/client"
	"ivar-go/src/models"
	"log"
	"net/http"
)

func UsersController(w http.ResponseWriter, _ *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error in GetFollowersController: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	ctx := context.Background()

	firestore, err := client.GetFirestoreClient()
	if err != nil {
		return
	}
	defer firestore.Close()

	iter := firestore.Collection("users").Documents(ctx)
	var users []models.User
	for {
		userSnap, err := iter.Next()
		var user models.User
		if err == iterator.Done {
			break
		}
		if err != nil {
			return
		}

		err = userSnap.DataTo(&user)
		if err != nil {
			return
		}

		user.ID = userSnap.Ref.ID
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		return
	}
}
