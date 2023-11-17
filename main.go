package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fish895623/bilf/handlers"
	"github.com/fish895623/bilf/route"
	"github.com/fish895623/bilf/templates"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupRouter() (e *gin.Engine) {
	e = gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	return
}

type CustomEngine struct {
	e *gin.Engine
}

func (r CustomEngine) Routing() {
	g := r.e.Group("/")
	g.GET("/", func(c *gin.Context) {
		var body templates.String
		body.Str = `<h1>hello world</h1>`
		body.Str += fmt.Sprintf(`<a href="https://google.com">%v</a>`, "google")
		body.Header()
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(body.Str))
	})
	g.GET("/metrics", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"tags": 1})
	})
	g.GET("/index/:id", func(c *gin.Context) {
		userid := c.Param("id")
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Hello?" + userid})
	})
}

// NOTE Query about gorm https://gorm.io/docs/query.html
func main() {
	var err error
	if err = godotenv.Load(".env"); err != nil {
		log.Println("Could not load .env file")
	}

	gin.SetMode(gin.ReleaseMode)
	e := SetupRouter()

	CustomEngine{e}.Routing()
	auth := e.Group("/auth")
	auth.Use(handlers.DummyMiddleWare)
	auth.Any("/auth", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
		case "POST":
			c.JSON(http.StatusOK, gin.H{"asdf": "post"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"asdf": "asdf"})
	})
	route.RouterRoot(e, "/")
	route.AuthRoute(e)
	e.Run()
}
