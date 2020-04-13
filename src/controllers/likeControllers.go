package controllers

import (
	"encoding/json"
	"ivar-go/src/client"
	"ivar-go/src/impl"
	"ivar-go/src/models"
	"log"
	"net/http"
)

func AddLikeToPost(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error in AddLikeController: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	fc, err := client.GetFirestoreClient()
	if err != nil {
		return
	}

	defer fc.Close()

	var addLikeBody models.AddLike

	_ = json.NewDecoder(r.Body).Decode(&addLikeBody)

	err = impl.AddLikeToPost(fc, addLikeBody)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetLikersForPost(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error in GetLikersForPostController: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	fc, err := client.GetFirestoreClient()
	if err != nil {
		return
	}

	defer fc.Close()

	queries := r.URL.Query()
	username := queries["u"][0]
	postId := queries["p"][0]

	likersData, err := impl.GetLikers(fc, username, postId)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(likersData)
	if err != nil {
		return
	}
}
