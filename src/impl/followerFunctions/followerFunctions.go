package followerFunctions

import (
	"cloud.google.com/go/firestore"
	"context"
	"ivar-go/src/mapper"
	"ivar-go/src/models"
)

func GetFollowers(fc *firestore.Client, username string) ([]models.GetFollowersResponse, error) {
	usersSnap, err := fc.Collection("users").Doc(username).Get(context.Background())
	if err != nil {
		return []models.GetFollowersResponse{}, err
	}

	var followerRefs models.FollowerRefs

	err = usersSnap.DataTo(&followerRefs)
	if err != nil {
		return []models.GetFollowersResponse{}, err
	}

	followersSnaps, err := fc.GetAll(context.Background(), followerRefs.FollowersRefs)
	if err != nil {
		return []models.GetFollowersResponse{}, err
	}

	var followersData []models.GetFollowersResponse

	for _, fs := range followersSnaps {
		var followerData models.User
		var followerResponse models.GetFollowersResponse

		err = fs.DataTo(&followerData)
		if err != nil {
			return []models.GetFollowersResponse{}, err
		}

		followerResponse = mapper.UserToFollowerResponse(followerData, followerResponse)
		followerResponse.Username = fs.Ref.ID

		followersData = append(followersData, followerResponse)
	}

	return followersData, nil
}
