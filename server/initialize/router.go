package initialize

import (
	"gin/middleware"
	"gin/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	Router.Use(middleware.Cors())

	systemRouter := router.GroupApp.System
	platformRouter := router.GroupApp.Platform

	PrivateGroup := Router.Group("/api")

	{
		systemRouter.InitApiRouter(PrivateGroup)      //注册功能api路由
		systemRouter.InitLoginApiRouter(PrivateGroup) //注册登录api路由
		systemRouter.InitUploadRouter(PrivateGroup)   //注册上传api路由
		systemRouter.InitUserRouter(PrivateGroup)     //注册系统用户api路由
	}
	{
		platformRouter.InitBlockchainRouter(PrivateGroup)    //注册区块链信息api路由
		platformRouter.InitAssetTransferRouter(PrivateGroup) // 注册asset-transfer链码路由
	}

	return Router
}
