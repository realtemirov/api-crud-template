package models

type Post struct {
	Base
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
	UserID  int64  `json:"user_id" db:"user_id"`
}

type GetAllPostsResponse struct {
	Meta Meta   `json:"meta"`
	Data []Post `json:"data"`
}
