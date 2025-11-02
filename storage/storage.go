package storage

import (
	"github.com/Javohir070/medium/storage/mysql"
	"github.com/Javohir070/medium/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	User() repo.UserStorageI
	Post() repo.PostStorageI
}

type storageMysql struct {
	userRepo repo.UserStorageI
	postRepo repo.PostStorageI
}

func NewStorage(mysqlConn *sqlx.DB) StorageI {
	return &storageMysql{
		userRepo: mysql.NewUserStorage(mysqlConn),
		postRepo: mysql.NewPostStorage(mysqlConn),
	}
}

func (s *storageMysql) User() repo.UserStorageI {
	return s.userRepo
}

func (s *storageMysql) Post() repo.PostStorageI {
	return s.postRepo
}
