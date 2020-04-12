package models

import (
	"cloud.google.com/go/firestore"
	"time"
)

type User struct {
	CreatedAt time.Time                `firestore:"createdAt"`
	Email     string                   `firestore:"email"`
	FirstName string                   `firestore:"firstName"`
	Followers []*firestore.DocumentRef `firestor:"followers"`
	Following []*firestore.DocumentRef `firestor:"following"`
	LastName  string                   `firestore:"lastName"`
	UpdatedAt time.Time                `firestore:"updatedAt"`
}

type GetUserResponse struct {
	CreatedAt      time.Time `firestore:"createdAt"`
	Email          string    `firestore:"email"`
	FirstName      string    `firestore:"firstName"`
	FollowerCount  int       `firestore:"followerCount"`
	FollowingCount int       `firestore:"followingCount"`
	LastName       string    `firestore:"lastName"`
	Posts          []GetPost `firestore:"posts"`
	UpdatedAt      time.Time `firestore:"updatedAt"`
	Username       string    `firestore:"username"`
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
	Text      string    `firestore:"text"`
	ImageUrl  string    `firestore:"imageUrl"`
	Likes     []string  `firestore:"likes"`
	CreatedAt time.Time `firestore:"createdAt"`
	UpdatedAt time.Time `firestore:"updatedAt"`
	Comments  []Comment `firestore:"comments"`
}

type GetPost struct {
	ID        string    `firestore:"id"`
	Text      string    `firestore:"text"`
	ImageUrl  string    `firestore:"imageUrl"`
	Likes     []string  `firestore:"likes"`
	CreatedAt time.Time `firestore:"createdAt"`
	UpdatedAt time.Time `firestore:"updatedAt"`
	Comments  []Comment `firestore:"comments"`
}

type CreatePost struct {
	Text     string `firestore:"id"`
	ImageUrl string `firestore:"id"`
}

type Comment struct {
	ID        string    `firestore:"id"`
	Username  string    `firestore:"username"`
	Likes     []string  `firestore:"likes"`
	CreatedAt time.Time `firestore:"createdAt"`
	UpdatedAt time.Time `firestore:"updatedAt"`
	Comments  []Comment `firestore:"comments"`
}

type FollowerRefs struct {
	FollowersRefs []*firestore.DocumentRef `firestore:"followers"`
}
