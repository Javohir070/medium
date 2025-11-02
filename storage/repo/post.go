package repo

import (
	"context"
	"time"
)

type PostStorageI interface {
	Create(ctx context.Context, req *Post) (*Post, error)
	Get(ctx context.Context, id string) (*Post, error)
	Delete(ctx context.Context, id string) error
}

type Post struct {
	ID        int       `db:"id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	UserID    string    `db:"user_id"`
	Published bool      `db:"published"`
	CreatedAt time.Time `db:"created_at"`
}
