package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/pkg/api/handler"
	"main.go/pkg/api/middleware"
)

func Loginroute(engine *gin.RouterGroup, loginhandler *handler.LoginHNadler) {
	engine.POST("/login", loginhandler.Login)
	engine.POST("/signup", loginhandler.Signup)

	//middileware
	//engine.Use(middleware.AuthMiddleware)
	restricted := engine.Group("/restricted")
	restricted.Use(middleware.AuthMiddleware)
	{
		restricted.GET("", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Hello, you are authenticated!")
		})
	}
}
