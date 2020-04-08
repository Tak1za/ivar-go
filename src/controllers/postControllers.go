package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"ivar-go/src/client"
	"ivar-go/src/models"
	"log"
	"net/http"
)

func GetPostsByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctx := context.Background()

	firestore, _ := client.GetFirestoreClient()
	defer firestore.Close()

	doc, err := firestore.Collection("users").Doc(vars["userId"]).Get(ctx)
	if err != nil {
		log.Fatalf("Error fetching data from firestore: %s", err)
	}

	postsData, _ := doc.DataAt("posts")

	var posts []models.Post

	jsonString, _ := json.Marshal(postsData)
	err = json.Unmarshal(jsonString, &posts)
	if err != nil {
		log.Fatalf("Error unmarshalling to json: %s", err)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		log.Fatalf("Error encoding data: %s", err)
	}
}
