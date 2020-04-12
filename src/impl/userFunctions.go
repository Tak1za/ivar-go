package impl

import (
	"cloud.google.com/go/firestore"
	"context"
	"ivar-go/src/mapper"
	"ivar-go/src/models"
	"time"
)

func GetUser(fc *firestore.Client, username string) (models.GetUserResponse, error) {
	userSnap, err := fc.Collection("users").Doc(username).Get(context.Background())
	if err != nil {
		return models.GetUserResponse{}, err
	}

	var user models.User
	var userData models.GetUserResponse

	err = userSnap.DataTo(&user)
	if err != nil {
		return models.GetUserResponse{}, err
	}

	//Get User Posts
	userPosts, err := GetPosts(fc, username)
	if err != nil {
		return models.GetUserResponse{}, nil
	}

	userData = mapper.UserToGetUserResponse(user, userData)
	userData.Username = userSnap.Ref.ID
	userData.Posts = userPosts

	return userData, nil
}

func CreateUser(fc *firestore.Client, createUserBody models.CreateUser) error {
	var newUser models.User

	newUser.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newUser.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newUser.Email = createUserBody.Email
	newUser.FirstName = createUserBody.FirstName
	newUser.LastName = createUserBody.LastName
	newUser.Followers = []*firestore.DocumentRef{}
	newUser.Following = []*firestore.DocumentRef{}

	_, err := fc.Collection("users").Doc(createUserBody.Username).Set(context.Background(), newUser)
	if err != nil {
		return err
	}

	return nil
}
