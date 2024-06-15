package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/pkg/api/handler"
	"main.go/pkg/api/middleware"
)

func Loginroute(engine *gin.RouterGroup, loginhandler *handler.LoginHNadler) {
	engine.POST("/login", loginhandler.Login)
	engine.POST("/signup", loginhandler.Signup)

	engine.Use(middleware.AuthMiddleware)
	{
		restricted := engine.Group("/restricted")
		{
			restricted.POST("/add", loginhandler.MakeToDO)
			restricted.PATCH("/update", loginhandler.UpdateToDo)
			restricted.DELETE("/remove", loginhandler.DeleteToDo)
		}
	}
}
