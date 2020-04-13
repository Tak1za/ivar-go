package impl

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"ivar-go/src/mapper"
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

func GetLikers(fc *firestore.Client, username string, postId string) ([]models.GetLikersResponse, error) {
	path := fmt.Sprintf("users/%s/posts/%s", username, postId)
	postSnap, err := fc.Doc(path).Get(context.Background())
	if err != nil {
		return []models.GetLikersResponse{}, err
	}

	var likerRefs models.LikerRefs

	err = postSnap.DataTo(&likerRefs)
	if err != nil {
		return []models.GetLikersResponse{}, err
	}

	likersSnaps, err := fc.GetAll(context.Background(), likerRefs.LikerRefs)
	if err != nil {
		return []models.GetLikersResponse{}, err
	}

	var likersData []models.GetLikersResponse

	for _, ls := range likersSnaps {
		var likerData models.User
		var likersResponse models.GetLikersResponse

		err = ls.DataTo(&likerData)
		if err != nil {
			return []models.GetLikersResponse{}, err
		}

		likersResponse = mapper.UserToLikerResponse(likerData, likersResponse)
		likersResponse.Username = ls.Ref.ID

		likersData = append(likersData, likersResponse)
	}

	return likersData, nil
}
