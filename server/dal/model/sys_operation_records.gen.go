// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysOperationRecord = "sys_operation_records"

// SysOperationRecord 操作记录表
type SysOperationRecord struct {
	ID           SnowflakeID    `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deletedAt"`
	IP           string         `gorm:"column:ip;type:varchar(255);comment:请求ip" json:"ip"`               // 请求ip
	Method       string         `gorm:"column:method;type:varchar(255);comment:请求方法" json:"method"`       // 请求方法
	Path         string         `gorm:"column:path;type:varchar(255);comment:请求路径" json:"path"`           // 请求路径
	Status       int32          `gorm:"column:status;type:int;comment:请求状态" json:"status"`                // 请求状态
	Latency      string         `gorm:"column:latency;type:varchar(255);comment:延迟" json:"latency"`       // 延迟
	Agent        string         `gorm:"column:agent;type:text;comment:代理" json:"agent"`                   // 代理
	ErrorMessage string         `gorm:"column:error_message;type:text;comment:错误信息" json:"errorMessage"`  // 错误信息
	Body         string         `gorm:"column:body;type:text;comment:请求Body" json:"body"`                 // 请求Body
	Resp         string         `gorm:"column:resp;type:text;comment:响应Body" json:"resp"`                 // 响应Body
	UserID       SnowflakeID    `gorm:"column:user_id;type:bigint;primaryKey;comment:用户id" json:"userId"` // 用户id
}

// TableName SysOperationRecord's table name
func (*SysOperationRecord) TableName() string {
	return TableNameSysOperationRecord
}
