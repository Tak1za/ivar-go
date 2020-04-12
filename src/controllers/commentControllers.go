package controllers

import (
	"encoding/json"
	"ivar-go/src/client"
	"ivar-go/src/impl"
	"ivar-go/src/models"
	"log"
	"net/http"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error in CreateCommentController: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	fc, err := client.GetFirestoreClient()
	if err != nil {
		return
	}

	defer fc.Close()

	var createCommentBody models.CreateComment

	_ = json.NewDecoder(r.Body).Decode(&createCommentBody)

	createdCommentId, err := impl.CreateComment(fc, createCommentBody)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdCommentId)
	if err != nil {
		return
	}
}
