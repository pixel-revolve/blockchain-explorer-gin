package system

import (
	api "gin/api"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("api")
	apiRouterApi := api.GroupApp.SystemApiGroup.SystemApiApi

	{
		apiRouter.GET("testApi", apiRouterApi.TestApi)                                   // 测试api
		apiRouter.GET("testParamApi", apiRouterApi.TestParamApi)                         // 测试问好传参api
		apiRouter.POST("testPostFormApi", apiRouterApi.TestPostFormApi)                  // 测试表单传参api
		apiRouter.POST("testJsonApi", apiRouterApi.TestJsonApi)                          // 测试application/json传参api
		apiRouter.POST("testRestfulApi/:username/:address", apiRouterApi.TestRestfulApi) // 测试restful传参api
	}
}
