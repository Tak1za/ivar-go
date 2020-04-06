package elastic_client

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func GetESClient(elasticUrl string, username string, password string) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			elasticUrl,
		},
		Username: username,
		Password: password,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error initializing the client: %s", err)
	}

	return es, nil
}
