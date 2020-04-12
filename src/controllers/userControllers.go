package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"ivar-go/src/client"
	"ivar-go/src/impl/userFunctions"
	"ivar-go/src/models"
	"log"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error in GetUserController: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	fc, err := client.GetFirestoreClient()
	if err != nil {
		return
	}
	defer fc.Close()

	vars := mux.Vars(r)
	username := vars["username"]

	userData, err := userFunctions.GetUser(fc, username)
	if err != nil {
		return
	}

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
			log.Printf("Error in CreateUserController: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	fc, err := client.GetFirestoreClient()
	if err != nil {
		return
	}
	defer fc.Close()

	var createUserBody models.CreateUser

	_ = json.NewDecoder(r.Body).Decode(&createUserBody)

	err = userFunctions.CreateUser(fc, createUserBody)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
}
