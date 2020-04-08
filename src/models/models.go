package models

type User struct {
	ID        string   `json:"id,omitempty"`
	CreatedAt string   `json:"createdAt,omitempty"`
	Email     string   `json:"email,omitempty"`
	FirstName string   `json:"firstName,omitempty"`
	Friends   []string `json:"friends,omitempty"`
	LastName  string   `json:"lastName,omitempty"`
	Posts     []Post   `json:"posts,omitempty"`
	UpdatedAt string   `json:"updatedAt,omitempty"`
	Username  string   `json:"username,omitempty"`
}

type Post struct {
	ID        string    `json:"id,omitempty"`
	Text      string    `json:"text,omitempty"`
	ImageUrl  string    `json:"imageUrl,omitempty"`
	Likes     []string  `json:"likes,omitempty"`
	Username  string    `json:"username,omitempty"`
	CreatedAt string    `json:"createdAt,omitempty"`
	UpdatedAt string    `json:"updatedAt,omitempty"`
	Comments  []Comment `json:"comments,omitempty"`
}

type Comment struct {
	ID        string    `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Likes     []string  `json:"likes,omitempty"`
	CreatedAt string    `json:"createdAt,omitempty"`
	UpdatedAt string    `json:"updatedAt,omitempty"`
	Comments  []Comment `json:"comments,omitempty"`
}
