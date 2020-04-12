package controllers

import (
	firestore2 "cloud.google.com/go/firestore"
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	var err error
	var errNotFound error
	defer func() {
		if err != nil {
			log.Printf("Error in GetFollowersController: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		if errNotFound != nil {
			log.Printf("Error in GetPostByPostIdController: %v", err)
			w.WriteHeader(http.StatusNotFound)
		}
	}()

	firestore, err := client.GetFirestoreClient()
	if err != nil {
		return
	}
	defer firestore.Close()

	vars := mux.Vars(r)
	userId := vars["userId"]

	userSnap, errNotFound := firestore.Collection("users").Doc(userId).Get(context.Background())
	if errNotFound != nil {
		return
	}

	var user models.User
	var userData models.GetUserResponse

	err = userSnap.DataTo(&user)
	if err != nil {
		return
	}

	userData.Username = userSnap.Ref.ID
	userData.FollowerCount = len(user.Followers)
	userData.FollowingCount = len(user.Following)
	userData.LastName = user.LastName
	userData.FirstName = user.FirstName
	userData.Email = user.Email
	userData.CreatedAt = user.CreatedAt
	userData.UpdatedAt = user.UpdatedAt

	path := fmt.Sprintf("users/%s/posts", userId)
	iter := firestore.Collection(path).Documents(context.Background())
	var userPosts []models.GetPost
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
		userPosts = append(userPosts, post)
	}

	userData.Posts = userPosts

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(userData)
	if err != nil {
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error in GetFollowersController: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	firestore, err := client.GetFirestoreClient()
	if err != nil {
		return
	}
	defer firestore.Close()

	var createUser models.CreateUser
	var newUser models.User

	_ = json.NewDecoder(r.Body).Decode(&createUser)
	newUser.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newUser.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newUser.Email = createUser.Email
	newUser.FirstName = createUser.FirstName
	newUser.LastName = createUser.LastName
	newUser.Followers = []*firestore2.DocumentRef{}

	wr, err := firestore.Collection("users").Doc(createUser.Username).Set(context.Background(), newUser)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(wr)
	if err != nil {
		return
	}
}
