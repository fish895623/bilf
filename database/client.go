package database

import (
	"fmt"
	"os"

	"github.com/fish895623/bilf/types"
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

func init() {
	var err error

	if DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{}); err != nil {
		panic("Failed to connect to database")
	}
	DB.AutoMigrate(&types.Tag{}, &types.Daily{})
}
