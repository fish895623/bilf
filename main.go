package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fish895623/bilf/route"
	T "github.com/fish895623/bilf/types"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "1234"
	DB_NAME     = "hhhh"
)

var DBINFO = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
var db *gorm.DB

func DBInit() {
	var err error

	db, err = gorm.Open(postgres.Open(DBINFO), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	db.AutoMigrate(&T.Daily{})
	db.AutoMigrate(&T.Tag{})
}

func SetupRouter() (e *gin.Engine) {
	e = gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.LoadHTMLGlob("templates/*.html")
	return
}

func main() {
	gin.SetMode(gin.DebugMode)
	e := SetupRouter()
	DBInit()

	route.RouterRoot(e, "/")

	e.GET("/index/:id", func(c *gin.Context) {
		userid := c.Param("id")
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Hello?" + userid})
	})
	e.Run()
}
