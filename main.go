package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "1234"
	DB_NAME     = "hhhh"
)

var dbinfo = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
var db *sql.DB

type Settings struct {
	Database string
}
type Tag struct {
	Id   int
	Name string
}
type Daily struct {
	Tag []int
}

func route_root(e *gin.Engine, address string) (g *gin.RouterGroup) {
	g = e.Group(address)
	g.GET("", func(c *gin.Context) {
		var db *sql.DB
		var err error
		if db, err = sql.Open("postgres", dbinfo); err != nil {
			panic(err.Error())
		}
		var rows *sql.Rows
		rows, err = db.Query("SELECT * FROM taglist")
		defer db.Close()

		var id int
		var name string
		var context_TagList []Tag
		for rows.Next() {
			var tag Tag
			if err := rows.Scan(&id, &name); err != nil {
				panic(err.Error())
			}
			tag.Id = id
			tag.Name = name
			context_TagList = append(context_TagList, tag)
		}
		for _, tag := range context_TagList {
			fmt.Println(tag.Id, tag.Name)
		}

		if err != nil {
			panic(err.Error())
		}

		c.HTML(http.StatusOK, "index.html", gin.H{"taglist": context_TagList})
	})

	return
}

func Database() (db *sql.DB) {
	return
}

func SetupRouter() (e *gin.Engine) {
	e = gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.LoadHTMLGlob("templates/*.html")
	return
}

func main() {

	// NOTE. only postgresql?
	var err error
	if db, err = sql.Open("postgres", dbinfo); err != nil {
		panic(err.Error())
	}
	defer db.Close()

	gin.SetMode(gin.DebugMode)
	e := SetupRouter()

	route_root(e, "/")

	e.GET("/index/:id", func(c *gin.Context) {
		userid := c.Param("id")
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Hello?" + userid})
	})
	e.Run()
}
