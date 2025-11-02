package mysql

import (
	"context"
	"database/sql"

	"github.com/Javohir070/medium/storage/repo"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserStorage(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(ctx context.Context, req *repo.User) (*repo.User, error) {
	query := `INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)`
	result, err := u.db.ExecContext(ctx, query, req.FirstName, req.LastName, req.Email, req.Password)
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

func (u *userRepo) Update(ctx context.Context, req *repo.UpdateUser) error {
	tsx, err := u.db.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE users SET
			first_name = ?,
			last_name = ?
		WHERE id = ? 
	`
	res, err := tsx.Exec(query, req.FirstName, req.LastName, req.ID)

	if err != nil {
		errRoll := tsx.Rollback()
		if errRoll != nil {
			err = errRoll
		}
		return err
	}
	data, err := res.RowsAffected()

	if err != nil {
		errRoll := tsx.Rollback()
		if errRoll != nil {
			err = errRoll
		}
		return err
	}
	if data == 0 {
		tsx.Commit()
		return sql.ErrNoRows
	}

	return tsx.Commit()
}

func (u *userRepo) Get(ctx context.Context, id string) (*repo.User, error) {
	query := `SELECT id, first_name, last_name, email, password, created_at FROM users WHERE id = ?`
	var user repo.User
	err := u.db.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := u.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
