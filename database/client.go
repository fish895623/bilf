package database

import (
	"fmt"
	"log"

	T "github.com/fish895623/bilf/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "1234"
	DB_NAME     = "hhhh"
)

var DBINFO = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
var DB *gorm.DB

func DBInit() {
	var err error

	DB, err = gorm.Open(postgres.Open(DBINFO), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	DB.AutoMigrate(&T.Daily{})
	DB.AutoMigrate(&T.Tag{})
	return
}
