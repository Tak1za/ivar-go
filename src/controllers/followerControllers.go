package controllers

import (
	firestore2 "cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
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

	vars := mux.Vars(r)

	firestore, err := client.GetFirestoreClient()
	if err != nil {
		return
	}

	defer firestore.Close()

	usersRef, errNotFound := firestore.Collection("users").Doc(vars["userId"]).Get(context.Background())
	if errNotFound != nil {
		return
	}

	followersRef, err := usersRef.DataAt("followers")
	if followersRef == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		return
	}

	//Doing this because we are storing userRef inside followers array
	var followersRefs []*firestore2.DocumentRef

	jsonString, _ := json.Marshal(followersRef)
	err = json.Unmarshal(jsonString, &followersRefs)
	if err != nil {
		return
	}

	followersSnaps, errNotFound := firestore.GetAll(context.Background(), followersRefs)
	if errNotFound != nil {
		return
	}

	var followersData []models.User

	for _, fs := range followersSnaps {
		var followerData models.User

		jsonString, _ = json.Marshal(fs.Data())
		err = json.Unmarshal(jsonString, &followerData)
		if err != nil {
			return
		}

		followersData = append(followersData, followerData)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(followersData)
	if err != nil {
		return
	}
}
