package models

type CreatePost struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserID    string `json:"user_id"`
	Published bool   `json:"published"`
}

type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserID    string `json:"user_id"`
	Published bool   `json:"published"`
}

type UpdatePost struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	Published bool   `json:"published"`
}
