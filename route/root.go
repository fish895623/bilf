package route

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupMiddleWare(e *gin.Engine) {
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
}

func Setup() (e *gin.Engine) {
	e = gin.New()
	SetupMiddleWare(e)
	Routing(e)

	return
}

func Routing(e *gin.Engine) {
	g := e.Group("/")
	g.GET("/", CookieTool())
	g.GET("/login", SetCookies())
}

func SetCookies() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("label", "ok", 3600, "/", "localhost", true, true)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`<a>Login</a>`))
	}
}
func CookieTool() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("label"); err == nil {
			if cookie == "ok" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden with no cookie"})
		c.Abort()
	}
}

type Price struct {
	High float64 `json:"high"`
	Low  float64 `json:"low"`
}

type ResponseData struct {
	Status              string  `json:"status"`
	Today               Price   `json:"today"`
	Week                Price   `json:"week"`
	Year                Price   `json:"year"`
	DividendsPercentage float64 `json:"DividendsPercentage"`
}
