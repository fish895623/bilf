package route

import (
	"context"
	"log"
	"net/http"

	"github.com/chromedp/chromedp"
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
		ctx, cancel := chromedp.NewContext(
			context.Background(),
		)
		defer cancel()

		// navigate to a page
		var example string
		err := chromedp.Run(ctx,
			chromedp.Navigate(`https://finance.yahoo.com/quote/JEPI`),
			chromedp.Text(`td[data-test="VTD_DTR-value"]`, &example),
		)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(example)

		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
