package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fish895623/bilf/route"
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

func DummyMiddleWare(c *gin.Context) {
	fmt.Println("DummyMiddleWare")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Cache-Control", "no-cache, no-cache, must-revalidate, post-check=0, pre-check=0, max-age=0")
	c.Header("Last-Modified", time.Now().String())
	c.Header("Expires", "-1")
	c.Next()
}

// NOTE Query about gorm https://gorm.io/docs/query.html

func main() {
	gin.SetMode(gin.ReleaseMode)
	e := SetupRouter()

	e.Use(DummyMiddleWare)
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
