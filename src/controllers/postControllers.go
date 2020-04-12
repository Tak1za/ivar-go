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
	"time"
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

	postsRef := firestore.Collection("users").Doc(vars["userId"]).Collection("posts")
	iter := postsRef.Documents(context.Background())
	defer iter.Stop()

	var posts []models.GetPost

	for {
		var post models.GetPost
		postSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if postSnap == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err != nil {
			return
		}

		err = postSnap.DataTo(&post)
		if err != nil {
			return
		}

		post.ID = postSnap.Ref.ID
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
	postSnap, errNotFound := firestore.Collection(path).Doc(vars["postId"]).Get(context.Background())
	if errNotFound != nil {
		return
	}

	var post models.GetPost

	err = postSnap.DataTo(&post)
	if err != nil {
		return
	}
	post.ID = postSnap.Ref.ID

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
			log.Printf("Error in GetPostByPostIdController: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	vars := mux.Vars(r)

	firestore, err := client.GetFirestoreClient()
	if err != nil {
		return
	}
	defer firestore.Close()

	var createPost models.CreatePost
	var newPost models.Post

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&createPost)
	newPost.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newPost.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newPost.Text = createPost.Text
	newPost.ImageUrl = createPost.ImageUrl
	newPost.Comments = []models.Comment{}
	newPost.Likes = []string{}

	path := fmt.Sprintf("users/%s/posts", vars["userId"])
	createdPost, _, err := firestore.Collection(path).Add(context.Background(), newPost)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdPost.ID)
	if err != nil {
		return
	}
}
