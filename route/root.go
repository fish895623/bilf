package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func RouterRoot(e *gin.Engine, address string) (g *gin.RouterGroup) {
	g = e.Group(address)
	g.GET("", tagListDisplay)
	g.POST("taglist", createNewTag)

	return
}
func tagListDisplay(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"taglist": ""})
}

func createNewTag(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"tags": 1})
}
