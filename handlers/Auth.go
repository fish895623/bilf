package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func DummyMiddleWare(c *gin.Context) {
	// requiredToken := os.Getenv("API_TOKEN")
	requiredToken := "123456"

	if requiredToken == "" {
		log.Fatal("env API_TOKEN")
	}

	token := c.Request.Header.Get("api_token")

	if token == requiredToken {
		return
	}
	if token != requiredToken {
		return
	}
	c.Next()
}
