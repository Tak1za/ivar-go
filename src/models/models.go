package models

import (
	"cloud.google.com/go/firestore"
	"time"
)

type User struct {
	ID        string    `firestore:"id"`
	CreatedAt time.Time `firestore:"createdAt"`
	Email     string    `firestore:"email"`
	FirstName string    `firestore:"firstName"`
	Friends   []string  `firestore:"friends"`
	LastName  string    `firestore:"lastName"`
	Posts     []Post    `firestore:"posts"`
	UpdatedAt time.Time `firestore:"updatedAt"`
	Username  string    `firestore:"username"`
}

type Post struct {
	ID        string    `firestore:"id"`
	Text      string    `firestore:"text"`
	ImageUrl  string    `firestore:"imageUrl"`
	Likes     []string  `firestore:"likes"`
	Username  string    `firestore:"username"`
	CreatedAt time.Time `firestore:"createdAt"`
	UpdatedAt time.Time `firestore:"updatedAt"`
	Comments  []Comment `firestore:"comments"`
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
