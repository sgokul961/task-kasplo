package apis

import (
	"github.com/gin-gonic/gin"
	"main.go/pkg/api/handler"
	"main.go/pkg/routes"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHttp(handler *handler.LoginHNadler) *ServerHTTP {
	engine := gin.New()
	engine.Use(gin.Logger())
	routes.Loginroute(engine.Group(""), handler)
	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
