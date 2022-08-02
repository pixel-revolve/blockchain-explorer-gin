package system

import "gin/service"

// ApiGroup 同一个包下的内容可以直接调用
type ApiGroup struct {
	SystemApiApi
	LoginApi
	UploadApi
	BaseApi
}

var (
	userService = service.GroupApp.SystemServiceGroup.UserService
)
