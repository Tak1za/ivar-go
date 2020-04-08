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
	router.HandleFunc("/users", controllers.UsersController)
	router.HandleFunc("/posts/{userId}", controllers.GetPostsByUserId)
	router.HandleFunc("/post/{userId}/{postId}", controllers.GetPostByPostId)
	fmt.Println("IVAR-Go listening at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
