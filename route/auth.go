package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRoute(e *gin.Engine) (g *gin.RouterGroup) {
	g = e.Group("/auth")
	g.Any("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "auth/login.html", gin.H{"status": "ok"})
	})

	return
}
