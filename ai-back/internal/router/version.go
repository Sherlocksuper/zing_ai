package router

import (
	"awesomeProject3/api"
	"awesomeProject3/internal/handler"
	"awesomeProject3/internal/service"
	"github.com/gin-gonic/gin"
)

var versionService = service.NewVersionService()
var versionHandler = handler.NewVersionHandler(versionService)

func RegisterVersionRoute(router *gin.Engine) {

	versionGroup := router.Group(api.API + "/version")
	{
		versionGroup.GET("/all", versionHandler.GetAllVersions)
		versionGroup.POST("/add", versionHandler.AddVersion)
		versionGroup.GET("/latest", versionHandler.GetLatestVersion)
		versionGroup.POST("/update", versionHandler.UpdateVersion)
	}
}
