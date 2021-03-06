package controllers

import (
	"encoding/json"
	"ivar-go/src/client"
	"ivar-go/src/impl"
	"log"
	"net/http"
)

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error in GetFollowersController: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	queries := r.URL.Query()
	username := queries["u"][0]

	fc, err := client.GetFirestoreClient()
	if err != nil {
		return
	}

	defer fc.Close()

	followersData, err := impl.GetFollowers(fc, username)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(followersData)
	if err != nil {
		return
	}
}
