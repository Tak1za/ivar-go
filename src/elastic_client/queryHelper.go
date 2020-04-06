package elastic_client

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"ivar-go/src/helpers"
	"log"
)

func SearchQuery(queryBody bytes.Buffer) []byte {
	index := helpers.Config.Database.Index
	elasticUrl := helpers.Config.Server.URL
	username := helpers.Config.Server.Username
	password := helpers.Config.Server.Password

	es, _ := GetESClient(elasticUrl, username, password)

	resp, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
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
		log.Fatalf("Error reading the response body: %s", err)
	}

	return respBody
}
