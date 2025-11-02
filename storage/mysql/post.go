package mysql

import (
	"context"

	"github.com/Javohir070/medium/storage/repo"
	"github.com/jmoiron/sqlx"
)

type postRepo struct {
	db *sqlx.DB
}

func NewPostStorage(db *sqlx.DB) repo.PostStorageI {
	return &postRepo{db: db}
}

func (p *postRepo) Create(ctx context.Context, req *repo.Post) (*repo.Post, error) {
	query := `INSERT INTO posts (title, body, user_id, published) VALUES (?,  ?, ?, ?)`
	result, err := p.db.ExecContext(ctx, query, req.Title, req.Body, req.UserID, req.Published)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	req.ID = int(id)
	return req, nil
}

func (p *postRepo) Get(ctx context.Context, id string) (*repo.Post, error) {
	query := `SELECT id, title, body, user_id, published FROM posts WHERE id = ?`
	var post repo.Post
	err := p.db.GetContext(ctx, &post, query, id)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (p *postRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM posts WHERE id = ?`

	_, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
