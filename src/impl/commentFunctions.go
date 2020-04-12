package impl

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"ivar-go/src/models"
	"time"
)

func CreateComment(fc *firestore.Client, createCommentBody models.CreateComment) (string, error) {
	var newComment models.Comment

	newComment.Username = createCommentBody.CurrentUser
	newComment.Text = createCommentBody.Text
	newComment.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newComment.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newComment.Likes = []*firestore.DocumentRef{}
	newComment.ParentComment = createCommentBody.ParentComment

	path := fmt.Sprintf("users/%s/posts/%s/comments", createCommentBody.PostOwner, createCommentBody.PostId)
	createdComment, _, err := fc.Collection(path).Add(context.Background(), newComment)
	if err != nil {
		return "", err
	}

	return createdComment.ID, nil
}
