package route

import (
	_ "fmt"
	"net/http"

	"github.com/fish895623/bilf/database"
	"github.com/fish895623/bilf/types"
	"github.com/gin-gonic/gin"
)

func RouterRoot(e *gin.Engine, address string) (g *gin.RouterGroup) {
	g = e.Group(address)
	g.GET("taglist", func(c *gin.Context) {
		var as []types.Tag

		database.DB.Find(&as)
		c.JSON(http.StatusOK, gin.H{"data": as})
	})

	return
}

func createNewTag(c *gin.Context) {
}
