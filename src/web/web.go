package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"ivar-go/src/controllers"
	"ivar-go/src/helpers"
	"log"
	"net/http"
)

func main() {
	//Configuration Setup
	helpers.ReadConfiguration()

	//Router Setup
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.HomeController)
	router.HandleFunc("/users", controllers.UsersController)
	fmt.Println("IVAR-Go listening at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
