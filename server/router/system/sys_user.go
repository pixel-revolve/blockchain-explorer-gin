package system

import (
	api "gin/api"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	userRouterWithoutRecord := Router.Group("user")
	baseApi := api.GroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("register", baseApi.Register)            // 注册账号
		userRouter.DELETE("deleteUser", baseApi.DeleteUser)      // 删除用户
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)       // 设置用户信息
		userRouter.PUT("changePassword", baseApi.ChangePassword) /// 修改用户密码
	}
	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList)            // 分页获取用户列表
		userRouterWithoutRecord.GET("findUserById/:id", baseApi.FindUserById)       // 获取自身信息
		userRouterWithoutRecord.GET("findUserByUUID/:uuid", baseApi.FindUserByUUID) // 通过uuid获取自身信息
	}
}
