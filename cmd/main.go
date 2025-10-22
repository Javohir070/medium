package main

import (
	"fmt"
	"log"

	"github.com/Javohir070/medium/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := config.Load(".")

	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.Mysql.User,
		cfg.Mysql.Password,
		cfg.Mysql.Host,
		cfg.Mysql.Port,
		cfg.Mysql.DBName,
	)

	mysqlConn, err := sqlx.Connect("mysql", mysqlUrl)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Successfully connected to the database")
	fmt.Println("Database Connection Info:")
	fmt.Println(mysqlUrl)

	_ = mysqlConn
}
