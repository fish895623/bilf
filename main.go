package main

import (
	"fmt"
	"net/http"

	"github.com/fish895623/bilf/database"
	"github.com/fish895623/bilf/route"
	"github.com/fish895623/bilf/types"
	"github.com/gin-gonic/gin"
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

// NOTE Query about gorm https://gorm.io/docs/query.html

func main() {
	gin.SetMode(gin.DebugMode)
	e := SetupRouter()
	var as []types.Tag

	database.DB.Find(&as)
	fmt.Printf("%+v#n", as)

	e.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	e.GET("/metrics", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"tags": 1})
	})

	e.GET("/index/:id", func(c *gin.Context) {
		userid := c.Param("id")
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Hello?" + userid})
	})
	route.RouterRoot(e, "/")

	e.Run()
}
