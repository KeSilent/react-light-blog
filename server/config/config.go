package config

import "github.com/kesilent/react-light-blog/config/types"

type Server struct {
	System types.System `mapstructure:"system" json:"system" yaml:"system"`
	JWT    types.JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    types.Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`

	//gorm
	Mysql  types.Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Sqlite types.Sqlite `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	// oss
	Local types.Local `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu types.Qiniu `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	// 跨域配置
	Cors types.CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
