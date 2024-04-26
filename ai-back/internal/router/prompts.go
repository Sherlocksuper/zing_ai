package router

import (
	"awesomeProject3/api"
	"awesomeProject3/internal/handler"
	"awesomeProject3/internal/service"
	"github.com/gin-gonic/gin"
)

var promptService = service.NewPromptService()

var promptHandler = handler.NewPromptHandler(promptService)

func RegisterPromptsRouter(router *gin.Engine) {

	promptGroup := router.Group(api.API + "/prompt")
	{
		promptGroup.POST("/add", promptHandler.AddPrompt)
		promptGroup.GET("/delete", promptHandler.DeletePrompt)
		promptGroup.GET("/list", promptHandler.GetPromptList)
		promptGroup.POST("/update", promptHandler.UpdatePrompt)
	}
}
