package system

import (
	"context"
	"errors"

	"github.com/kesilent/react-light-blog/dal/model"
	"github.com/kesilent/react-light-blog/dal/query"
	systemReq "github.com/kesilent/react-light-blog/dal/request"
	"github.com/kesilent/react-light-blog/utils"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type UserService struct{}

var UserServiceApp = new(UserService)

// @author: JackYang
// @function: Register
// @description: 用户注册
// @param: u model.SysUser
// @return: userInter model.SysUser, err error
func (userService *UserService) Register(u model.SysUser) (userInter model.SysUser, err error) {
	q := query.Q.SysUser

	_, err = q.WithContext(context.Background()).Where(q.Username.Eq(u.Username)).First()

	if !errors.Is(err, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.ID, err = utils.GenID(0)
	if err != nil {
		return userInter, err
	}
	err = q.WithContext(context.Background()).Save(&u)
	return u, err
}

// @author: JackYang
// @function: Login
// @description: 登陆
// @param: u model.SysUser
// @return: userInter model.SysUser, err error
func (userService *UserService) Login(u *model.SysUser) (userInter *model.SysUser, err error) {
	var user *model.SysUser
	q := query.Q.SysUser
	user, err = q.WithContext(context.Background()).Where(q.Username.Eq(u.Username)).Preload(field.NewRelation("Authorities", "")).First()
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		MenuServiceApp.UserAuthorityDefaultRouter(&user.Authorities[0])
	}
	return user, err
}

// @author: JackYang
// @function: ChangePassword
// @description: 修改用户密码
// @param: u *model.SysUser, newPassword string
// @return: userInter *model.SysUser,err error
func (userService *UserService) ChangePassword(u *model.SysUser, newPassword string) (userInter *model.SysUser, err error) {
	var user *model.SysUser
	q := query.Q.SysUser
	user, err = q.WithContext(context.Background()).Where(q.ID.Eq(u.ID)).First()
	if err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = q.WithContext(context.Background()).Save(user)
	return user, err

}

// @author: JackYang
// @function: GetUserInfoList
// @description: 分页获取数据
// @param: info request.PageInfo
// @return: err error, list interface{}, total int64
func (userService *UserService) GetUserInfoList(info systemReq.GetUserList) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := query.Q.SysUser.WithContext(context.Background())

	if info.NickName != "" {
		db = db.Where(query.SysUser.NickName.Like("%" + info.NickName + "%"))
	}
	if info.Phone != "" {
		db = db.Where(query.SysUser.Phone.Like("%" + info.Phone + "%"))
	}
	if info.Username != "" {
		db = db.Where(query.SysUser.Username.Like("%" + info.Username + "%"))
	}
	if info.Email != "" {
		db = db.Where(query.SysUser.Email.Like("%" + info.Email + "%"))
	}

	total, err = db.Count()
	if err != nil {
		return
	}

	userList, err := db.Preload(field.NewRelation("Authorities", "")).Preload(field.NewRelation("Authority", "")).Limit(limit).Offset(offset).Find()

	return userList, total, err
}

// @author: JackYang
// @function: SysUser
// @description: 通过ID获取用户信息
// @param: id int64
// @return: *model.SysUser, error
func (userService *UserService) GetUserInfo(id int64) (*model.SysUser, error) {
	q := query.Q.SysUser
	return q.WithContext(context.Background()).Preload(field.NewRelation("Authorities", "")).Where(query.SysUser.ID.Eq(id)).First()
}

func (userService *UserService) AddUserAuthorities(userAuthorities []*model.SysUserAuthority) error {
	q := query.Q.SysUserAuthority
	return q.WithContext(context.Background()).Create(userAuthorities...)
}
