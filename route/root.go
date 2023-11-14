package route

import (
	"database/sql"
	"net/http"

	DB "github.com/fish895623/bilf/db"
	I "github.com/fish895623/bilf/types"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func RouterRoot(e *gin.Engine, address string) (g *gin.RouterGroup) {
	g = e.Group(address)
	g.GET("", func(c *gin.Context) {
		var db *sql.DB
		var err error
		if db, err = sql.Open("postgres", DB.DBINFO); err != nil {
			panic(err.Error())
		}
		var rows *sql.Rows
		rows, err = db.Query("SELECT * FROM taglist")
		defer db.Close()

		var id int
		var name string
		var context_TagList []I.Tag
		for rows.Next() {
			var tag I.Tag
			if err := rows.Scan(&id, &name); err != nil {
				panic(err.Error())
			}
			tag.Id = id
			tag.Name = name
			context_TagList = append(context_TagList, tag)
		}

		if err != nil {
			panic(err.Error())
		}

		c.HTML(http.StatusOK, "index.html", gin.H{"taglist": context_TagList})
	})

	return
}
