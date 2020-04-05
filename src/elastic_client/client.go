package elastic_client

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func GetESClient() (*elasticsearch.Client, error) {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error initializing the client: %s", err)
	}

	return es, nil
}
