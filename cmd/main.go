package main

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"go.mod/cmd/api"
	"go.mod/cmd/db"
)

func main() {
	db, error := db.NewMysqlStorage(mysql.Config{
		User:   "root",
		Passwd: "root",
		Addr:   "127.0.1:3306",
		DBName: "expenseapp",
	})

	if error != nil {
		log.Fatal(error)
	} else {
		err := db.Ping()
		if err != nil {
			log.Fatal("There was a problem connecting to the database: ?", err)
		} else {
			fmt.Println("Connection to the database was successful")
		}
	}

	server := api.NewApiServer("localhost:8000", db)
	server.Run()
}
