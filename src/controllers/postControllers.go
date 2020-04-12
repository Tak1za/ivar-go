package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"ivar-go/src/client"
	"ivar-go/src/impl/postFunctions"
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

	username := r.Header.Get("username")

	fc, err := client.GetFirestoreClient()
	if err != nil {
		return
	}

	defer fc.Close()

	posts, err := postFunctions.GetPosts(fc, username)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		return
	}
}

func GetPostByPostId(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error in GetPostByPostIdController: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	vars := mux.Vars(r)
	postId := vars["postId"]
	username := r.Header.Get("username")

	fc, err := client.GetFirestoreClient()
	if err != nil {
		return
	}
	defer fc.Close()

	post, err := postFunctions.GetPost(fc, username, postId)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		return
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error in CreatePostController: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	username := r.Header.Get("username")

	fc, err := client.GetFirestoreClient()
	if err != nil {
		return
	}
	defer fc.Close()

	var createPostBody models.CreatePost

	_ = json.NewDecoder(r.Body).Decode(&createPostBody)

	createdPostId, err := postFunctions.CreatePost(fc, username, createPostBody)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdPostId)
	if err != nil {
		return
	}
}
