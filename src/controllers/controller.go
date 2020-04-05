package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	es, _ := elastic_client.GetESClient()

	queryBody := queries.GetUsersQuery()
	resp, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("application"),
		es.Search.WithBody(&queryBody),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer resp.Body.Close()

	if resp.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			log.Fatalf("[%s] %s: %s",
				resp.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var elasticResponse models.ESResponse

	err = json.Unmarshal(respBody, &elasticResponse)
	if err != nil {
		log.Fatalln(err)
	}

	responseData := elasticResponse.OuterHits.InnerHits
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		log.Fatalln(err)
	}
}
