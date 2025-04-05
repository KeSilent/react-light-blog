// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysUser = "sys_users"

// SysUser 用户表
type SysUser struct {
	ID         int64      `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true;comment:用户ID" json:"id"` // 用户ID
	UUID       string     `gorm:"column:uuid;type:char(36);not null" json:"uuid"`
	Username   string     `gorm:"column:username;type:varchar(64);not null;comment:用户名" json:"username"`       // 用户名
	Password   string     `gorm:"column:password;type:varchar(255);not null;comment:密码" json:"password"`       // 密码
	NickName   string     `gorm:"column:nick_name;type:varchar(64);comment:昵称" json:"nickName"`                // 昵称
	Avatar     string     `gorm:"column:avatar;type:varchar(255);comment:头像" json:"avatar"`                    // 头像
	Phone      string     `gorm:"column:phone;type:varchar(11);comment:手机号" json:"phone"`                      // 手机号
	Email      string     `gorm:"column:email;type:varchar(128);comment:邮箱" json:"email"`                      // 邮箱
	Status     bool       `gorm:"column:status;type:tinyint(1);default:1;comment:状态(0:禁用,1:启用)" json:"status"` // 状态(0:禁用,1:启用)
	CreateTime time.Time  `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP" json:"createTime"`
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`
	DeletedAt  *time.Time `gorm:"column:deleted_at;type:datetime" json:"deletedAt"`
	Role       []SysRole  `gorm:"foreignKey:uuid;joinForeignKey:sys_user_uuid;joinReferences:sys_role_uuid;many2many:sys_user_role;references:uuid" json:"role"`
	Dept       []SysDept  `gorm:"foreignKey:uuid;joinForeignKey:sys_user_uuid;joinReferences:sys_dept_uuid;many2many:sys_user_dept;references:uuid" json:"dept"`
}

// TableName SysUser's table name
func (*SysUser) TableName() string {
	return TableNameSysUser
}
