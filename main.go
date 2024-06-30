package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Users struct {
	Firstname  string         `db:"firstname"`
	Middlename sql.NullString `db:"middlename"`
	Lastname   string         `db:"lastname"`
}

func main() {
	db, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/expenseapp")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("database connected successfully\n")
	}

	var users = []Users{}

	err = db.Select(&users, "SELECT firstname,middlename,lastname FROM user_info")
	if err != nil {
		fmt.Println("b")
		fmt.Println(err)
		return
	} else {
		fmt.Println(users)
	}

	defer db.Close()
}
