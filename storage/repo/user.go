package repo

import (
	"context"
	"time"
)

type UserStorageI interface {
	Create(ctx context.Context, req *User) (*User, error)
	Update(ctx context.Context, req *UpdateUser) error
	Get(ctx context.Context, id string) (*User, error)
	Delete(ctx context.Context, id string) error
}

type User struct {
	ID        int       `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type UpdateUser struct {
	ID        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}
