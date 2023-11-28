package route

import (
	"encoding/json"
	"fmt"
	"io"
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
		req, err := http.Get("https://finance.yahoo.com/quote/JEPII")
		if err != nil {
			log.Fatal(err.Error())
			// TODO send error message when failed to connect
			c.JSON(http.StatusOK, gin.H{
				"status":       "failed",
				"ErrorMessage": "Fail to get html", "log": err.Error(),
			})
		}
		defer req.Body.Close()

		html, err := goquery.NewDocumentFromReader(req.Body)
		if err != nil {
			log.Fatal(err.Error())
		}

		c.JSON(http.StatusOK, gin.H{
			"status":              "ok",
			"DividendsPercentage": html.Find(`td[data-test="YTD_DTR-value"]`).Text(),
		})
	})
	g.POST("/ping", func(c *gin.Context) {
		body := c.Request.Body
		val, err := io.ReadAll(body)
		if err != nil {
			log.Fatal(err.Error())
		}
		var data struct {
			Name string `json:"name"`
		}

		json.Unmarshal([]byte(val), &data)
		log.Println(data.Name)

		req, err := http.Get(fmt.Sprintf(`https://finance.yahoo.com/quote/%s`, data.Name))
		if err != nil {
			log.Fatal(err.Error())
			// TODO send error message when failed to connect
			c.JSON(http.StatusOK, gin.H{
				"status":       "failed",
				"ErrorMessage": "Fail to get html", "log": err.Error(),
			})
		}
		defer req.Body.Close()

		html, err := goquery.NewDocumentFromReader(req.Body)
		if err != nil {
			log.Fatal(err.Error())
		}

		c.JSON(http.StatusOK, gin.H{
			"status":              "ok",
			"DividendsPercentage": html.Find(`td[data-test="YTD_DTR-value"]`).Text(),
		})
	})
}
