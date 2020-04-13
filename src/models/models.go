package models

import (
	"cloud.google.com/go/firestore"
	"time"
)

type User struct {
	CreatedAt time.Time                `firestore:"createdAt"`
	Email     string                   `firestore:"email"`
	FirstName string                   `firestore:"firstName"`
	Followers []*firestore.DocumentRef `firestore:"followers"`
	Following []*firestore.DocumentRef `firestore:"following"`
	LastName  string                   `firestore:"lastName"`
	UpdatedAt time.Time                `firestore:"updatedAt"`
}

type GetUserResponse struct {
	CreatedAt      time.Time         `firestore:"createdAt"`
	Email          string            `firestore:"email"`
	FirstName      string            `firestore:"firstName"`
	FollowerCount  int               `firestore:"followerCount"`
	FollowingCount int               `firestore:"followingCount"`
	LastName       string            `firestore:"lastName"`
	Posts          []GetPostResponse `firestore:"posts"`
	UpdatedAt      time.Time         `firestore:"updatedAt"`
	Username       string            `firestore:"username"`
}

type GetLikersResponse struct {
	FirstName string `firestore:"firstName"`
	LastName  string `firestore:"lastName"`
	Username  string `firestore:"username"`
}

type GetFollowersResponse struct {
	FirstName string `firestore:"firstName"`
	LastName  string `firestore:"lastName"`
	Username  string `firestore:"username"`
}

type CreateUser struct {
	FirstName string `firestore:"firstName"`
	LastName  string `firestore:"lastName"`
	Email     string `firestore:"email"`
	Username  string `firestore:"username"`
}

type Post struct {
	Text      string                   `firestore:"text"`
	ImageUrl  string                   `firestore:"imageUrl"`
	Likes     []*firestore.DocumentRef `firestore:"likes"`
	CreatedAt time.Time                `firestore:"createdAt"`
	UpdatedAt time.Time                `firestore:"updatedAt"`
}

type GetPostResponse struct {
	ID         string    `firestore:"id"`
	Text       string    `firestore:"text"`
	ImageUrl   string    `firestore:"imageUrl"`
	LikesCount int       `firestore:"likesCount"`
	CreatedAt  time.Time `firestore:"createdAt"`
	UpdatedAt  time.Time `firestore:"updatedAt"`
}

type CreatePost struct {
	Username string `firestore:"username"`
	Text     string `firestore:"id"`
	ImageUrl string `firestore:"imageUrl"`
}

type Comment struct {
	Username      string                   `firestore:"username"`
	Text          string                   `firestore:"text"`
	Likes         []*firestore.DocumentRef `firestore:"likes"`
	CreatedAt     time.Time                `firestore:"createdAt"`
	UpdatedAt     time.Time                `firestore:"updatedAt"`
	ParentComment string                   `firestore:"parentComment"`
}

type CreateComment struct {
	CurrentUser   string `firestore:"currentUser"`
	PostId        string `firestore:"postId"`
	PostOwner     string `firestore:"postOwner"`
	Text          string `firestore:"text"`
	ParentComment string `firestore:"parentComment"`
}

type AddLike struct {
	CurrentUser string `firestore:"currentUser"`
	EntityId    string `firestore:"entityId"`
	EntityOwner string `firestore:"entityOwner"`
}

type FollowerRefs struct {
	FollowersRefs []*firestore.DocumentRef `firestore:"followers"`
}

type LikerRefs struct {
	LikerRefs []*firestore.DocumentRef `firestore:"likes"`
}
