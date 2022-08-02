package system

import (
	"gin/api"
	"github.com/gin-gonic/gin"
)

type LoginApiRouter struct{}

func (l *LoginApiRouter) InitLoginApiRouter(Router *gin.RouterGroup) {
	loginRouter := Router.Group("api/login")

	loginRouterApi := api.GroupApp.SystemApiGroup.LoginApi

	{
		loginRouter.POST("loginJson", loginRouterApi.LoginJson)
	}

}
