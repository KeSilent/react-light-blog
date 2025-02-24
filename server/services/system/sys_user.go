package system

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/kesilent/react-light-blog/global"
	"github.com/kesilent/react-light-blog/models/system"
	systemReq "github.com/kesilent/react-light-blog/models/system/request"
	"github.com/kesilent/react-light-blog/utils"
	"gorm.io/gorm"
)

type UserService struct{}

var UserServiceApp = new(UserService)

// @author: JackYang
// @function: Register
// @description: 用户注册
// @param: u model.SysUser
// @return: userInter system.SysUser, err error
func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.RLB_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.New()
	err = global.RLB_DB.Create(&u).Error
	return u, err
}

func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.RLB_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser
	err = global.RLB_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		MenuServiceApp.UserAuthorityDefaultRouter(&user)
	}
	return &user, err
}

// @author: JackYang
// @function: ChangePassword
// @description: 修改用户密码
// @param: u *model.SysUser, newPassword string
// @return: userInter *model.SysUser,err error
func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (userInter *system.SysUser, err error) {
	var user system.SysUser
	if err = global.RLB_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.RLB_DB.Save(&user).Error
	return &user, err

}

// @author: JackYang
// @function: GetUserInfoList
// @description: 分页获取数据
// @param: info request.PageInfo
// @return: err error, list interface{}, total int64
func (userService *UserService) GetUserInfoList(info systemReq.GetUserList) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.RLB_DB.Model(&system.SysUser{})
	var userList []system.SysUser

	if info.NickName != "" {
		db = db.Where("nick_name LIKE ?", "%"+info.NickName+"%")
	}
	if info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	if info.Email != "" {
		db = db.Where("email LIKE ?", "%"+info.Email+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}

// @author: JackYang
// @function: SetUserAuthority
// @description: 设置一个用户的权限
// @param: uuid uuid.UUID, authorityId string
// @return: err error
func (userService *UserService) SetUserAuthority(id uint, authorityId uint) (err error) {

	assignErr := global.RLB_DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&system.SysUserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}

	var authority system.SysAuthority
	err = global.RLB_DB.Where("authority_id = ?", authorityId).First(&authority).Error
	if err != nil {
		return err
	}
	var authorityMenu []system.SysAuthorityMenu
	var authorityMenuIDs []string
	err = global.RLB_DB.Where("sys_authority_authority_id = ?", authorityId).Find(&authorityMenu).Error
	if err != nil {
		return err
	}

	for i := range authorityMenu {
		authorityMenuIDs = append(authorityMenuIDs, authorityMenu[i].MenuId)
	}

	var authorityMenus []system.SysBaseMenu
	err = global.RLB_DB.Preload("Parameters").Where("id in (?)", authorityMenuIDs).Find(&authorityMenus).Error
	if err != nil {
		return err
	}
	hasMenu := false
	for i := range authorityMenus {
		if authorityMenus[i].Name == authority.DefaultRouter {
			hasMenu = true
			break
		}
	}
	if !hasMenu {
		return errors.New("找不到默认路由,无法切换本角色")
	}

	err = global.RLB_DB.Model(&system.SysUser{}).Where("id = ?", id).Update("authority_id", authorityId).Error
	return err
}
