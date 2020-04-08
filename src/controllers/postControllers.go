package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"ivar-go/src/client"
	"ivar-go/src/models"
	"log"
	"net/http"
)

func GetPostsByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	firestore, _ := client.GetFirestoreClient()
	defer firestore.Close()

	doc := firestore.Collection("users").Doc(vars["userId"]).Collection("posts")
	iter := doc.Documents(context.Background())
	defer iter.Stop()

	var posts []models.Post

	for {
		var post models.Post
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error iterating: %s", err)
		}

		jsonString, _ := json.Marshal(doc.Data())
		err = json.Unmarshal(jsonString, &post)
		if err != nil {
			log.Fatalf("Error unmarshalling to json: %s", err)
		}
		post.ID = doc.Ref.ID
		posts = append(posts, post)
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		log.Fatalf("Error encoding data: %s", err)
	}
}

func GetPostByPostId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctx := context.Background()

	firestore, _ := client.GetFirestoreClient()
	defer firestore.Close()

	doc, err := firestore.Collection(fmt.Sprintf("users/%s/posts", vars["userId"])).Doc(vars["postId"]).Get(ctx)
	if err != nil {
		log.Fatalf("Unable to fetch data: %s", err)
	}

	var post models.Post

	jsonString, _ := json.Marshal(doc.Data())
	err = json.Unmarshal(jsonString, &post)
	if err != nil {
		log.Fatalf("Error unmarshalling to json: %s", err)
	}
	post.ID = doc.Ref.ID

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		log.Fatalf("Error encoding data: %s", err)
	}
}
