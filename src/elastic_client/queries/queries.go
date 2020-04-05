package queries

import (
	"bytes"
	"encoding/json"
	"log"
)

func GetUsersQuery() (queryBuf bytes.Buffer) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	if err := json.NewEncoder(&queryBuf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	return queryBuf
}
