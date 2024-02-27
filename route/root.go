package route

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type CustomEngine struct {
	E *gin.Engine
}

func SetupMiddleWare(e *gin.Engine) {
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
}
func SetupRouter(e *gin.Engine) {
	CustomEngine{E: e}.Routing()
}

func Setup() (e *gin.Engine) {
	e = gin.New()
	SetupMiddleWare(e)
	SetupRouter(e)

	return
}

func (r CustomEngine) Routing() {
	g := r.E.Group("/")
	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"asdf": "asdf"})
	})
	g.Any("/ping", ANYPingStocks)
}
func ANYPingStocks(c *gin.Context) {
	var req *http.Response
	var html *goquery.Document

	if c.Request.Method == "GET" {
		req = RequestingQuote("JEPI")
	} else if c.Request.Method == "POST" {
		body := c.Request.Body
		val, _ := io.ReadAll(body)
		var data struct {
			Name string `json:"name"`
		}
		json.Unmarshal([]byte(val), &data)

		req = RequestingQuote(data.Name)
	}
	html, _ = goquery.NewDocumentFromReader(req.Body)
	var data = html.Find(`td[data-test="YTD_DTR-value"]`).Text()
	realData, _ := strconv.ParseFloat(strings.Replace(data, "%", "", -1), 64)

	if data == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "invalid stock name",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":              "ok",
			"DividendsPercentage": realData,
		})
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

func RequestingQuote(name string) (res *http.Response) {
	res, _ = http.Get(fmt.Sprintf("https://finance.yahoo.com/quote/%s", name))
	return
}
