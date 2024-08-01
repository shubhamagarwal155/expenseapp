package main

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"go.mod/pkg/api"
	"go.mod/pkg/conf"
)

func main() {
	db, error := conf.NewMysqlStorage(mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 "127.0.1:3306",
		DBName:               "expenseapp",
		AllowNativePasswords: true,
	})

	if error != nil {
		log.Fatal(error)
	} else {
		err := db.Ping()
		if err != nil {
			log.Fatal("There was a problem connecting to the database:", err)
		} else {
			fmt.Println("Connection to the database was successful")
		}
	}

	server := api.NewApiServer(db)
	err := server.Run("localhost:8000")

	if err != nil {
		log.Fatal("Error starting server : ", err.Error())
	}
}
