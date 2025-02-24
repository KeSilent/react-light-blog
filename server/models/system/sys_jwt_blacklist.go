package system

import "github.com/kesilent/react-light-blog/global"

type JwtBlacklist struct {
	global.RLB_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
