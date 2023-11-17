package main

import (
	"log"
	"net/http"

	"github.com/fish895623/bilf/handlers"
	"github.com/fish895623/bilf/route"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupRouter() (e *gin.Engine) {
	e = gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	return
}

// NOTE Query about gorm https://gorm.io/docs/query.html
func main() {
	var err error
	if err = godotenv.Load(".env"); err != nil {
		log.Println("Could not load .env file")
	}

	gin.SetMode(gin.ReleaseMode)
	e := SetupRouter()
	route.CustomEngine{E: e}.Routing()
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
	route.AuthRoute(e)
	e.Run()
}
