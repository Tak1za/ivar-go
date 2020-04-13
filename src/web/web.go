package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"ivar-go/src/controllers"
	"log"
	"net/http"
)

func main() {
	//Router Setup
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware)

	//User related routes
	router.HandleFunc("/users/{username}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")

	//Post related routes
	router.HandleFunc("/posts", controllers.GetPostsByUserId).Queries("u", "{u}").Methods("GET")
	router.HandleFunc("/posts/{postId}", controllers.GetPostByPostId).Queries("u", "{u}").Methods("GET")
	router.HandleFunc("/posts", controllers.CreatePost).Methods("POST")

	//Comment related routes
	router.HandleFunc("/comments", controllers.CreateComment).Methods("POST")

	//Like related routes
	router.HandleFunc("/likes", controllers.AddLikeToPost).Methods("POST")
	router.HandleFunc("/likes", controllers.GetLikersForPost).Queries("u", "{u}").Queries("p", "{p}").Methods("GET")

	//Follower related routes
	router.HandleFunc("/followers", controllers.GetFollowers).Queries("u", "{u}").Methods("GET")

	fmt.Println("IVAR-Go listening at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST,OPTIONS")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
