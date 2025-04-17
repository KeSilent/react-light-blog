// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysBaseMenu = "sys_base_menus"

// SysBaseMenu 菜单表
type SysBaseMenu struct {
	ID         SnowflakeID    `gorm:"column:id;type:bigint;primaryKey;comment:菜单ID" json:"id"`              // 菜单ID
	MenuLevel  int32          `gorm:"column:menu_level;type:int;comment:菜单层级" json:"menuLevel"`             // 菜单层级
	ParentID   SnowflakeID    `gorm:"column:parent_id;type:bigint;not null;comment:父菜单ID" json:"parentId"`  // 父菜单ID
	Path       string         `gorm:"column:path;type:varchar(128);comment:路由path" json:"path"`             // 路由path
	Name       string         `gorm:"column:name;type:varchar(64);comment:路由name" json:"name"`              // 路由name
	Hidden     bool           `gorm:"column:hidden;type:tinyint(1);comment:是否隐藏" json:"hidden"`             // 是否隐藏
	Component  string         `gorm:"column:component;type:varchar(128);comment:对应前端文件路径" json:"component"` // 对应前端文件路径
	Sort       int32          `gorm:"column:sort;type:int;comment:排序" json:"sort"`                          // 排序
	Title      string         `gorm:"column:title;type:varchar(64);comment:菜单名" json:"title"`               // 菜单名
	Icon       string         `gorm:"column:icon;type:varchar(64);comment:图标" json:"icon"`                  // 图标
	CreateTime time.Time      `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP" json:"createTime"`
	UpdateTime time.Time      `gorm:"column:update_time;type:datetime" json:"updateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deletedAt"`
	Role       []SysRole      `gorm:"joinForeignKey:sys_base_menu_id;joinReferences:sys_role_id;many2many:sys_role_menus" json:"role"`
	Children   []SysBaseMenu  `gorm:"foreignKey:parent_id" json:"children"`
}

// TableName SysBaseMenu's table name
func (*SysBaseMenu) TableName() string {
	return TableNameSysBaseMenu
}
