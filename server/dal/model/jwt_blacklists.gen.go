// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameJwtBlacklist = "jwt_blacklists"

// JwtBlacklist mapped from table <jwt_blacklists>
type JwtBlacklist struct {
	ID        SnowflakeID    `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime(3)" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime(3)" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deletedAt"`
	Jwt       string         `gorm:"column:jwt;type:text;comment:jwt" json:"jwt"` // jwt
}

// TableName JwtBlacklist's table name
func (*JwtBlacklist) TableName() string {
	return TableNameJwtBlacklist
}
