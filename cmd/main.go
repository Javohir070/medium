package main

import (
	// "context"
	"fmt"
	"log"

	"github.com/Javohir070/medium/config"
	"github.com/Javohir070/medium/server"
	"github.com/Javohir070/medium/storage"
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

	strg := storage.NewStorage(mysqlConn)

	router := server.NewServer(server.Options{
		Strg: strg,
	})

	if err := router.Run(cfg.Port); err != nil {
		log.Fatalf("Error starting server: %v", cfg.Port, err)
	}

	fmt.Println("Server is running on port", cfg.Port)

	// _ = strg

	// user, err := strg.User().Create(context.TODO(), &repo.User{
	// 	FirstName: "Javohir",
	// 	LastName:  "Qayumov",
	// 	Email:     "javohir.doe@example.com",
	// 	Password:  "securepassword",
	// })
	// if err != nil {
	// 	log.Fatalf("Error creating user: %v", err)
	// }

	// user, err := strg.User().Get(context.TODO(), "1")
	// if err != nil {
	// 	log.Fatalf("Error getting user: %v", err)
	// }

	// err = strg.User().Delete(context.Background(), "1")
	// if err != nil {
	// 	log.Fatalf("Error deleting user: %v", err)
	// }

	// log.Printf("User created: %+v\n", user)
	// fmt.Println(user)

}
