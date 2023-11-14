package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "1234"
	DB_NAME     = "hhhh"
)

var DBINFO = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
var DB *sql.DB

func OpenDatabase() {
	var err error
	if DB, err = sql.Open("postgres", DBINFO); err != nil {
		panic(err.Error())
	}
	defer DB.Close()
}
