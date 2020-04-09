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
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error in GetPostsByUserIdController: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	vars := mux.Vars(r)

	firestore, err := client.GetFirestoreClient()
	if err != nil {
		return
	}

	defer firestore.Close()

	doc := firestore.Collection("users").Doc(vars["userId"]).Collection("posts")
	iter := doc.Documents(context.Background())
	defer iter.Stop()

	var posts []models.Post

	for {
		var post models.Post
		doc, err := iter.Next()
		if doc == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err == iterator.Done {
			break
		}
		if err != nil {
			return
		}

		jsonString, _ := json.Marshal(doc.Data())
		err = json.Unmarshal(jsonString, &post)
		if err != nil {
			return
		}
		post.ID = doc.Ref.ID
		posts = append(posts, post)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		return
	}
}

func GetPostByPostId(w http.ResponseWriter, r *http.Request) {
	var err error
	var errNotFound error
	defer func() {
		if err != nil {
			log.Printf("Error in GetPostByPostIdController: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		if errNotFound != nil {
			log.Printf("Error in GetPostByPostIdController: %v", err)
			w.WriteHeader(http.StatusNotFound)
		}
	}()

	vars := mux.Vars(r)

	firestore, err := client.GetFirestoreClient()
	if err != nil {
		return
	}
	defer firestore.Close()

	path := fmt.Sprintf("users/%s/posts", vars["userId"])
	postData, errNotFound := firestore.Collection(path).Doc(vars["postId"]).Get(context.Background())
	if errNotFound != nil {
		return
	}

	var post models.Post

	jsonString, _ := json.Marshal(postData.Data())
	err = json.Unmarshal(jsonString, &post)
	if err != nil {
		return
	}
	post.ID = postData.Ref.ID

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		return
	}
}
