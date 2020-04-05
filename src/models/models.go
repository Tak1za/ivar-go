package models

type User struct {
	Username  string   `json:"username"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	CreatedAt int64    `json:"createdAt"`
	UpdatedAt int64    `json:"updatedAt"`
	Friends   []string `json:"friends"`
	Posts     []Post   `json:"posts"`
}

type Post struct {
	ID        string    `json:"id"`
	Text      string    `json:"content"`
	ImageUrl  string    `json:"imageUrl"`
	Likes     []string  `json:"likes"`
	Username  string    `json:"username"`
	CreatedAt int64     `json:"createdAt"`
	UpdatedAt int64     `json:"updatedAt"`
	Comments  []Comment `json:"comments"`
}

type Comment struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Likes     []string  `json:"likes"`
	CreatedAt int64     `json:"createdAt"`
	UpdatedAt int64     `json:"updatedAt"`
	Comments  []Comment `json:"comments"`
}
