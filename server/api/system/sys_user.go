package system

import (
	"gin/global"
	"gin/model/common/request"
	"gin/model/common/response"
	"gin/model/system"
	systemReq "gin/model/system/request"
	systemRes "gin/model/system/response"
	"gin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	//"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type BaseApi struct {}

//
// Register
// @Description:
// @receiver b
// @param c
//
func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password}
	err, userReturn := userService.Register(*user)

	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

//
// GetUserList
// @Description:
// @receiver b
// @param c
//
func (b *BaseApi) GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := userService.GetUserInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

//
// FindUserById
// @Description:
// @receiver b
// @param c
//
func (b *BaseApi) FindUserById(c *gin.Context) {

	if id, err :=strconv.Atoi(c.Param("id"));err==nil{
		if err,ReqUser:=userService.FindUserById(id);err!=nil{
			global.GVA_LOG.Error("获取失败!",zap.Error(err))
			response.FailWithMessage("获取失败",c)
			return
		}else {
			global.GVA_LOG.Info("获取成功!",zap.Int("id", id))
			response.OkWithDetailed(gin.H{"userInfo":ReqUser},"获取成功",c)
			return
		}
	}
	
	response.FailWithMessage("类型转换错误",c)
}


//
// FindUserByUUID
// @Description:
// @receiver b
// @param c
//
func (b *BaseApi) FindUserByUUID(c *gin.Context) {

	//uuid不为空
	if uuid:=c.Param("uuid");uuid!=""{
		if err,ReqUser:=userService.FindUserByUuid(uuid);err!=nil{
			global.GVA_LOG.Error("获取失败!",zap.Error(err))
			response.FailWithMessage("获取失败",c)
			return
		}else {
			global.GVA_LOG.Info("获取成功!",zap.String("uuid", uuid))
			response.OkWithDetailed(gin.H{"userInfo":ReqUser},"获取成功",c)
			return
		}
	}

	response.FailWithMessage("uuid不能为空",c)
}

// DeleteUser @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "用户ID"
// @Success 200 {object} response.Response{msg=string} "删除用户"
// @Router /user/deleteUser [delete]
func (b *BaseApi) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)

	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//jwtId := utils.GetUserID(c)
	//if jwtId == uint(reqId.ID) {
	//	response.FailWithMessage("删除失败, 自杀失败", c)
	//	return
	//}
	if err := userService.DeleteUser(reqId.ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

//
// ChangePassword
// @Description: 修改用户密码，最少需要传username，password，newPassword
// @accept: application/json
// @receiver b
// @param c
//
func (b *BaseApi) ChangePassword(c *gin.Context) {
	var user systemReq.ChangePasswordStruct
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := &system.SysUser{Username: user.Username, Password: user.Password}
	if err, _ := userService.ChangePassword(u, user.NewPassword); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// SetUserInfo
// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysUser true "ID, 用户名, 昵称"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "设置用户信息"
// @Router /user/SetUserInfo [put]
func (b *BaseApi) SetUserInfo(c *gin.Context) {
	//这里传入的数据需要id，nickname,phone,email
	//也就是暂时可以修改的数据是以上的数据
	var user systemReq.ChangeUserInfo
	_ = c.ShouldBindJSON(&user)
	//这里原来的逻辑是通过jwt解析获取用户id的
	//user.ID = utils.GetUserID(c)
	if err := userService.SetUserInfo(system.SysUser{
		GVA_MODEL: global.GVA_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		Phone:     user.Phone,
		Email:     user.Email,
	}); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}