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
	router.HandleFunc("/users", controllers.UsersController).Methods("GET")
	router.HandleFunc("/posts/{userId}", controllers.GetPostsByUserId).Methods("GET")
	router.HandleFunc("/post/{userId}/{postId}", controllers.GetPostByPostId).Methods("GET")
	router.HandleFunc("/followers/{userId}", controllers.GetFollowers).Methods("GET")
	router.HandleFunc("/posts/{userId}", controllers.CreatePost).Methods("POST")
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
