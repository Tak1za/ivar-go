package impl

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"ivar-go/src/models"
)

func AddLikeToPost(fc *firestore.Client, addLikeBody models.AddLike) error {
	path := fmt.Sprintf("users/%s/posts/%s", addLikeBody.EntityOwner, addLikeBody.EntityId)

	userRef := fc.Collection("users").Doc(addLikeBody.CurrentUser)

	_, err := fc.Doc(path).Update(context.Background(), []firestore.Update{{Path: "likes", Value: firestore.ArrayUnion(userRef)}})
	if err != nil {
		return err
	}

	return nil
}
