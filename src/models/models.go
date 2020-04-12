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
	LastName  string                   `firestore:"lastName"`
	UpdatedAt time.Time                `firestore:"updatedAt"`
}

type GetUser struct {
	ID        string                   `firestore:"id"`
	CreatedAt time.Time                `firestore:"createdAt"`
	Email     string                   `firestore:"email"`
	FirstName string                   `firestore:"firstName"`
	Followers []*firestore.DocumentRef `firestore:"followers"`
	LastName  string                   `firestore:"lastName"`
	Posts     []Post                   `firestore:"posts"`
	UpdatedAt time.Time                `firestore:"updatedAt"`
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
