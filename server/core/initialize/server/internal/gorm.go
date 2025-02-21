package internal

import (
	"time"

	types "github.com/kesilent/react-light-blog/config/types"
	"github.com/kesilent/react-light-blog/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
// Author [SliverHorn](https://github.com/SliverHorn)
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	var general types.GeneralDB
	switch global.RLB_CONFIG.System.DbType {
	case "mysql":
		general = global.RLB_CONFIG.Mysql.GeneralDB
	case "sqlite":
		general = global.RLB_CONFIG.Sqlite.GeneralDB
	default:
		general = global.RLB_CONFIG.Mysql.GeneralDB
	}
	return &gorm.Config{
		Logger: logger.New(NewWriter(general), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      general.LogLevel(),
			Colorful:      true,
		}),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
