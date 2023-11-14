package main

import (
	"fmt"
	"net/http"

	db "github.com/fish895623/bilf/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "1234"
	DB_NAME     = "hhhh"
)

var DBINFO = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)

func SetupRouter() (e *gin.Engine) {
	e = gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.LoadHTMLGlob("templates/*.html")
	return
}

func SomeHandler(db *gorm.DB, fn func(*gin.Context)) gin.HandlerFunc {
	return gin.HandlerFunc(fn)
}

func main() {
	gin.SetMode(gin.DebugMode)
	// e := SetupRouter()
	e := gin.Default()
	db.DBInit()
	var db gorm.DB
	e.GET("", SomeHandler(&db, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}))

	e.GET("/index/:id", func(c *gin.Context) {
		userid := c.Param("id")
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Hello?" + userid})
	})
	e.Run()
}
