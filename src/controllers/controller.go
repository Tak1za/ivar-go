package controllers

import (
	"encoding/json"
	"fmt"
	"ivar-go/src/elastic_client"
	"ivar-go/src/elastic_client/queries"
	"ivar-go/src/models"
	"log"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Welcome to the home page!")
}

func UsersController(w http.ResponseWriter, r *http.Request) {
	index := "application"

	//Get the required get users query
	queryBody := queries.GetUsersQuery()

	//Get the elastic request response (in bytes)
	respBody := elastic_client.SearchQuery(index, queryBody)

	var elasticResponse models.ESResponse

	//Convert []bytes to struct
	err := json.Unmarshal(respBody, &elasticResponse)
	if err != nil {
		log.Fatalln(err)
	}

	responseData := elasticResponse.OuterHits.InnerHits

	//Encode the data into the response writer
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		log.Fatalln(err)
	}
}
