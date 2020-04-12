package controllers

import (
	firestore2 "cloud.google.com/go/firestore"
	"encoding/json"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"ivar-go/src/client"
	"ivar-go/src/models"
	"log"
	"net/http"
	"time"
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
	var users []models.GetUser
	for {
		userSnap, err := iter.Next()
		var user models.GetUser
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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error in GetFollowersController: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	firestore, err := client.GetFirestoreClient()
	if err != nil {
		return
	}
	defer firestore.Close()

	var createUser models.CreateUser
	var newUser models.User

	_ = json.NewDecoder(r.Body).Decode(&createUser)
	newUser.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newUser.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newUser.Email = createUser.Email
	newUser.FirstName = createUser.FirstName
	newUser.LastName = createUser.LastName
	newUser.Followers = []*firestore2.DocumentRef{}

	wr, err := firestore.Collection("users").Doc(createUser.Username).Set(context.Background(), newUser)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(wr)
	if err != nil {
		return
	}
}
