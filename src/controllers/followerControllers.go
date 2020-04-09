package controllers

import (
	firestore2 "cloud.google.com/go/firestore"
	"encoding/json"
	"github.com/gorilla/mux"
	"ivar-go/src/client"
	"ivar-go/src/helpers"
	"ivar-go/src/models"
	"log"
	"net/http"
)

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	firestore, _ := client.GetFirestoreClient()
	defer firestore.Close()

	usersRef := helpers.GetUserRef(firestore, vars["userId"])

	followersRef := helpers.GetFollowersRef(usersRef)

	//Doing this because we are storing userRef inside followers array
	var followersRefs []*firestore2.DocumentRef

	jsonString, _ := json.Marshal(followersRef)
	err := json.Unmarshal(jsonString, &followersRefs)
	if err != nil {
		log.Fatalf("Error unmarshalling to array of documentRefs: %s", err)
	}

	followersSnaps := helpers.GetAllData(firestore, followersRefs)

	var followersData []models.User

	for _, fs := range followersSnaps {
		var followerData models.User

		jsonString, _ = json.Marshal(fs.Data())
		err = json.Unmarshal(jsonString, &followerData)

		followersData = append(followersData, followerData)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(followersData)
	if err != nil {
		log.Fatalf("Error encoding data: %s", err)
	}
}
