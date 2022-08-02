package system

import (
	"gin/api"
	"github.com/gin-gonic/gin"
)

type UploadRouter struct{}

func (u *UploadRouter) InitUploadRouter(Router *gin.RouterGroup) {

	uploadRouter := Router.Group("api/upload")
	uploadRouterApi := api.GroupApp.SystemApiGroup.UploadApi

	{
		uploadRouter.POST("file", uploadRouterApi.Upload)
	}

}
