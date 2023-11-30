package route

import (
	"github.com/fish895623/bilf/handlers"
	"github.com/gin-gonic/gin"
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
	SetupMiddleWare(e)
	SetupRouter(e)

	return
}
