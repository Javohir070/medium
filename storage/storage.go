package storage

import (
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	
}

type StorageMysql struct {
	
}

func NewStorage(mysqlConn *sqlx.DB) StorageI {
	return &StorageMysql{

	}
}
