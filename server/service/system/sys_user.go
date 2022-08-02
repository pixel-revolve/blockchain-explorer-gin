package system

import (
	"errors"
	"gin/global"
	"gin/model/common/request"
	"gin/model/system"
	"gin/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)


type UserService struct{}

func (userService *UserService) Register(u system.SysUser) (err error, userInter system.SysUser) {
	var user system.SysUser
	//实际开发使用这个方式判断错误
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	// 调用*gorm.DB类型的Create方法，在数据库中创建记录
	err = global.GVA_DB.Create(&u).Error
	// golang的独特处理错误方式。
	return err, u
}

func (userService *UserService) GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	//在上层的逻辑我们写的page是查询第几页，然后pageSize是每页展示多少的记录
	//然后在逻辑层面我们需要计算一下偏移量
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	// toDo:这里的结构要修改
	//err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return err, userList, total
}

func (userService *UserService) FindUserById(id int) (err error, user *system.SysUser) {
	var u system.SysUser
	err = global.GVA_DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

func (userService *UserService) FindUserByUuid(uuid string) (err error, user *system.SysUser) {
	var u system.SysUser
	if err = global.GVA_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}

func (userService *UserService) DeleteUser(id int) (err error) {
	var user system.SysUser
	err = global.GVA_DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	//这里之后加上权限之后需要进行处理
	//err = global.GVA_DB.Delete(&[]system.SysUseAuthority{}, "sys_user_id = ?", id).Error
	return err
}

// ChangePassword
//修改用户密码
func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (err error, userInter *system.SysUser) {
	var user system.SysUser
	err = global.GVA_DB.Where("username = ?", u.Username).First(&user).Error
	if err != nil {
		return err, nil
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return errors.New("原密码错误"), nil
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.GVA_DB.Save(&user).Error
	return err, &user

}

// SetUserInfo
//修改用户的部分信息
func (userService *UserService) SetUserInfo(req system.SysUser) error {
	return global.GVA_DB.Updates(&req).Error
}