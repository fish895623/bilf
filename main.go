package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fish895623/bilf/route"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func DummyMiddleWare() gin.HandlerFunc {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro ao ler variaveis de ambiente")
	}
	requiredToken := os.Getenv("API_TOKEN")

	if requiredToken == "" {
		log.Fatal("Por favor, defina a variavel API_TOKEN")
	}

	return func(c *gin.Context) {
		token := c.Request.Header.Get("api_token")

		if token == requiredToken {
			return
		}
		if token != requiredToken {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Token invalido"})
			c.Abort()
			return
		}

		c.Next()
	}

}

// NOTE Query about gorm https://gorm.io/docs/query.html

func main() {
	gin.SetMode(gin.ReleaseMode)
	e := SetupRouter()

	e.Use(DummyMiddleWare())
	e.GET("", func(c *gin.Context) {
		c.SetCookie("Access-Token", "123", 3600, "/", "localhost", false, true)
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
