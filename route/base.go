package route

import (
	"log"

	"github.com/fish895623/bilf/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupMiddleWare(e *gin.Engine) {
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(handlers.DummyMiddleWare)
}
func SetupRouter(e *gin.Engine) {
	CustomEngine{E: e}.Routing()
}

func Setup() (e *gin.Engine) {
	e = gin.New()
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Could not load .env file")
	}
	SetupMiddleWare(e)
	SetupRouter(e)

	return
}
