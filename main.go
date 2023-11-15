package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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


var db *gorm.DB

func main() {
	gin.SetMode(gin.DebugMode)
	e := SetupRouter()
	var err error
	if db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{}); err != nil {
		panic("Failed to connect to database")
	}
	e.GET("", SomeHandler(db, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}))

	e.GET("/index/:id", func(c *gin.Context) {
		userid := c.Param("id")
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Hello?" + userid})
	})
	e.Run()
}
