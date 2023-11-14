package main

import (
	"net/http"

	"github.com/fish895623/bilf/route"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

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

	route.RouterRoot(e, "/")

	e.GET("/index/:id", func(c *gin.Context) {
		userid := c.Param("id")
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Hello?" + userid})
	})
	e.Run()
}
