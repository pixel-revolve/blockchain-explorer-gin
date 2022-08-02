package system

import (
	"gin/global"
	"github.com/satori/go.uuid"
)

// SysUser
//定义SysUser模型，绑定sys_users表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
//我们可以通过结构体标签指定映射。然后默认的orm规则是将golang中的模型按照驼峰命名的格式给转成数据库的蛇形命名模式格式
//在这里SysUser类型可以代表mysql sys_users表
type SysUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID      `json:"uuid" gorm:"comment:用户UUID"`                                                           // 用户UUID
	Username    string         `json:"userName" gorm:"comment:用户登录名"`                                                        // 用户登录名
	Password    string         `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	NickName    string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Phone       string         `json:"phone"  gorm:"comment:用户手机号"` // 用户手机号
	Email       string         `json:"email"  gorm:"comment:用户邮箱"`  // 用户邮箱
}

// TableName
//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (SysUser) TableName() string {
	//return "sys_users"
	return "sys_users"
}
