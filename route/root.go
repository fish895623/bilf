package route

import (
	"fmt"
	"net/http"

	"github.com/fish895623/bilf/database"
	"github.com/fish895623/bilf/templates"
	"github.com/fish895623/bilf/types"
	"github.com/gin-gonic/gin"
)

func (r CustomEngine) Routing() {
	g := r.E.Group("/")
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
	g.GET("taglist", func(c *gin.Context) {
		var as []types.Tag

		database.DB.Find(&as)
		c.JSON(http.StatusOK, gin.H{"data": as})
	})
}
