package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func NewMysqlStorage(conf mysql.Config) (*sql.DB, error) {
	db, error := sql.Open("mysql", conf.FormatDSN())

	return db, error
}
