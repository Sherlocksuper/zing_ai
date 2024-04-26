package router

import (
	"awesomeProject3/api"
	"awesomeProject3/ws"
	"github.com/gin-gonic/gin"
)

func ConfigRouter(router *gin.Engine) {
	//router.Use(middleware.AuthTokenCheck())

	RegisterUserRouter(router)
	RegisterChatRouter(router)
	RegisterVersionRoute(router)
	RegisterPromptsRouter(router)
	RegisterAuthRouter(router)

	//配置跨域
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	router.GET(api.API+"/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/ws", ws.Handler)
}
