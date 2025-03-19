package request

import (
	common "github.com/kesilent/react-light-blog/dal/common/request"
)

// Register User register structure
type Register struct {
	Username     string   `json:"userName" example:"用户名"`
	Password     string   `json:"passWord" example:"密码"`
	NickName     string   `json:"nickName" example:"昵称"`
	HeaderImg    string   `json:"headerImg" example:"头像链接"`
	AuthorityId  uint     `json:"authorityId" swaggertype:"string" example:"int 角色id"`
	Enable       int      `json:"enable" swaggertype:"string" example:"int 是否启用"`
	AuthorityIds []uint64 `json:"authorityIds" swaggertype:"string" example:"[]uint 角色id"`
	Phone        string   `json:"phone" example:"电话号码"`
	Email        string   `json:"email" example:"电子邮箱"`
}

// Login User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// ChangePasswordReq Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// SetUserAuth Modify user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // 角色ID
}

type GetUserList struct {
	common.PageInfo
	Username string `json:"username" form:"username"`
	NickName string `json:"nickName" form:"nickName"`
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
}
