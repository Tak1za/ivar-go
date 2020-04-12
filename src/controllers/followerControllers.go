package controllers

import (
	"context"
	"encoding/json"
	"ivar-go/src/client"
	"ivar-go/src/models"
	"log"
	"net/http"
)

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	var err error
	var errNotFound error
	defer func() {
		if err != nil {
			log.Printf("Error in GetFollowersController: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		if errNotFound != nil {
			log.Printf("Error in GetFollowersController: %v", err)
			w.WriteHeader(http.StatusNotFound)
		}
	}()

	userId := r.Header.Get("userId")

	firestore, err := client.GetFirestoreClient()
	if err != nil {
		return
	}

	defer firestore.Close()

	usersSnap, errNotFound := firestore.Collection("users").Doc(userId).Get(context.Background())
	if errNotFound != nil {
		return
	}

	var followerRefs models.FollowerRefs

	err = usersSnap.DataTo(&followerRefs)
	if err != nil {
		return
	}

	followersSnaps, errNotFound := firestore.GetAll(context.Background(), followerRefs.FollowersRefs)
	if errNotFound != nil {
		return
	}

	var followersData []models.User

	for _, fs := range followersSnaps {
		var followerData models.User

		err = fs.DataTo(&followerData)
		if err != nil {
			return
		}
		followerData.ID = fs.Ref.ID
		followersData = append(followersData, followerData)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(followersData)
	if err != nil {
		return
	}
}
