package helpers

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"log"
)

func GetUserRef(client *firestore.Client, userId string) *firestore.DocumentSnapshot {

	userRef, err := client.Collection("users").Doc(userId).Get(context.Background())
	if err != nil {
		log.Fatalf("Error fetching userRef for userId: %s | %s", userId, err)
	}

	return userRef
}

func GetFollowersRef(userRef *firestore.DocumentSnapshot) interface{} {
	followersData, err := userRef.DataAt("followers")
	if err != nil {
		log.Fatalf("Error fetching followers data: %s", err)
	}

	return followersData
}

func GetPostRef(client *firestore.Client, userId string, postId string) *firestore.DocumentSnapshot {

	path := fmt.Sprintf("users/%s/posts", userId)
	postRef, err := client.Collection(path).Doc(postId).Get(context.Background())
	if err != nil {
		log.Fatalf("Error fetching postRef for userId: %s, postId: %s | %s", userId, postId, err)
	}

	return postRef
}

func GetAllData(client *firestore.Client, docRefs []*firestore.DocumentRef) []*firestore.DocumentSnapshot {
	docSnaps, err := client.GetAll(context.Background(), docRefs)
	if err != nil {
		log.Fatalf("Error fetching all data: %s", err)
	}

	return docSnaps
}
