package router

import (
	"awesomeProject3/api"
	"awesomeProject3/internal/handler"
	"awesomeProject3/internal/service"
	"github.com/gin-gonic/gin"
)

var chatService = service.NewChatService()
var chatHandler = handler.NewChatHandler(chatService)

func RegisterChatRouter(router *gin.Engine) {
	chatGroup := router.Group(api.API + "/chat")
	{
		chatGroup.POST("/start", chatHandler.StartAChat)
		chatGroup.GET("/delete", chatHandler.DeleteChat)
		chatGroup.GET("/deleteall", chatHandler.DeleteAllChat)
		chatGroup.GET("/detail", chatHandler.GetChatDetail)
		chatGroup.GET("/list", chatHandler.GetChatList)
		chatGroup.POST("/send", chatHandler.SendMessage)
		chatGroup.POST("/sendforimage", chatHandler.SendForImage)
	}
}
