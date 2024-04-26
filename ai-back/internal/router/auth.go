package router

import (
	"awesomeProject3/api"
	"awesomeProject3/internal/handler"
	"awesomeProject3/internal/service"
	"github.com/gin-gonic/gin"
)

var authService = service.NewAuthService()
var authHandler = handler.NewAuthHandler(authService)

func RegisterAuthRouter(router *gin.Engine) {
	auth := router.Group(api.API + "/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
		auth.POST("/resetpassword", authHandler.ResetPassword)
	}
}
