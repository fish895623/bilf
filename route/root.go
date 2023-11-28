package route

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type CustomEngine struct {
	E *gin.Engine
}

func (r CustomEngine) Routing() {
	g := r.E.Group("/")
	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"asdf": "asdf"})
	})
	g.GET("/ping", func(c *gin.Context) {
		req, err := http.Get("https://finance.yahoo.com/quote/JEPI")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer req.Body.Close()

		html, err := goquery.NewDocumentFromReader(req.Body)
		if err != nil {
			log.Fatal(err.Error())
		}

		c.JSON(http.StatusOK, gin.H{"message": "pong", "DividendsPercentage": html.Find(`td[data-test="YTD_DTR-value"]`).Text()})
	})
}
