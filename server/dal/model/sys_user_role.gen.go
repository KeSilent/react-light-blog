// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysUserRole = "sys_user_role"

// SysUserRole mapped from table <sys_user_role>
type SysUserRole struct {
	SysUserID SnowflakeID `gorm:"column:sys_user_id;type:bigint;not null" json:"sysUserId"`
	SysRoleID SnowflakeID `gorm:"column:sys_role_id;type:bigint;not null" json:"sysRoleId"`
}

// TableName SysUserRole's table name
func (*SysUserRole) TableName() string {
	return TableNameSysUserRole
}
