package storage

import (
	"github.com/Javohir070/medium/storage/mysql"
	"github.com/Javohir070/medium/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	User() repo.UserStorageI
}

type storageMysql struct {
	userRepo repo.UserStorageI
}

func NewStorage(mysqlConn *sqlx.DB) StorageI {
	return &storageMysql{
		userRepo: mysql.NewUserStorage(mysqlConn),
	}
}

func (s *storageMysql) User() repo.UserStorageI {
	return s.userRepo
}
