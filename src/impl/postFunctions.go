package impl

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"ivar-go/src/models"
	"time"
)

func GetPosts(fc *firestore.Client, username string) ([]models.GetPostResponse, error) {
	path := fmt.Sprintf("users/%s/posts", username)

	var posts []models.GetPostResponse

	iter := fc.Collection(path).Documents(context.Background())
	for {
		var post models.GetPostResponse
		postSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return []models.GetPostResponse{}, err
		}

		err = postSnap.DataTo(&post)
		if err != nil {
			return []models.GetPostResponse{}, err
		}
		post.ID = postSnap.Ref.ID
		posts = append(posts, post)
	}

	return posts, nil
}

func GetPost(fc *firestore.Client, username string, postId string) (models.GetPostResponse, error) {
	path := fmt.Sprintf("users/%s/posts", username)

	postSnap, err := fc.Collection(path).Doc(postId).Get(context.Background())
	if err != nil {
		return models.GetPostResponse{}, err
	}

	var post models.GetPostResponse

	err = postSnap.DataTo(&post)
	if err != nil {
		return models.GetPostResponse{}, err
	}
	post.ID = postSnap.Ref.ID

	return post, nil
}

func CreatePost(fc *firestore.Client, createPostBody models.CreatePost) (string, error) {
	var newPost models.Post

	newPost.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newPost.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newPost.Text = createPostBody.Text
	newPost.ImageUrl = createPostBody.ImageUrl
	newPost.Likes = []string{}

	path := fmt.Sprintf("users/%s/posts", createPostBody.Username)
	createdPost, _, err := fc.Collection(path).Add(context.Background(), newPost)
	if err != nil {
		return "", err
	}

	return createdPost.ID, nil
}
