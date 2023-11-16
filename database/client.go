package database

import (
	"os"
	"fmt"

	"github.com/fish895623/bilf/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error

	if DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{}); err != nil {
		fmt.Errorf("Failed to connect to database %w", err.Error())
	}
	DB.AutoMigrate(&types.Tag{}, &types.Daily{})
}
