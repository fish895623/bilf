package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterRoot(e *gin.Engine, address string) (g *gin.RouterGroup) {
	g = e.Group(address)
	g.GET("taglist", createNewTag)

	return
}

func createNewTag(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"tags": 1})
}
