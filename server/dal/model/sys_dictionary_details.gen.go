// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysDictionaryDetail = "sys_dictionary_details"

// SysDictionaryDetail 字典值
type SysDictionaryDetail struct {
	ID              int64          `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt       time.Time     `gorm:"column:created_at;type:datetime(3)" json:"createdAt"`
	UpdatedAt       time.Time     `gorm:"column:updated_at;type:datetime(3)" json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deletedAt"`
	Label           string        `gorm:"column:label;type:varchar(191);comment:展示值" json:"label"`                           // 展示值
	Value           int64         `gorm:"column:value;type:bigint;comment:字典值" json:"value"`                                 // 字典值
	Status          int64         `gorm:"column:status;type:bigint;comment:启用状态" json:"status"`                              // 启用状态
	Sort            int64         `gorm:"column:sort;type:bigint;comment:排序标记" json:"sort"`                                  // 排序标记
	SysDictionaryID int64         `gorm:"column:sys_dictionary_id;type:bigint unsigned;comment:关联标记" json:"sysDictionaryId"` // 关联标记
	Dictionary      SysDictionary  `gorm:"foreignKey:sys_dictionary_id;references:id" json:"dictionary"`
}

// TableName SysDictionaryDetail's table name
func (*SysDictionaryDetail) TableName() string {
	return TableNameSysDictionaryDetail
}
