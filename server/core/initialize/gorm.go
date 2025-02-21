package initialize

import (
	"os"

	"github.com/kesilent/react-light-blog/core/initialize/server"
	"github.com/kesilent/react-light-blog/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.RLB_CONFIG.System.DbType {
	case "mysql":
		global.RLB_ACTIVE_DBNAME = &global.RLB_CONFIG.Mysql.Dbname
		return server.GormMysql()
	case "sqlite":
		global.RLB_ACTIVE_DBNAME = &global.RLB_CONFIG.Sqlite.Dbname
		return server.GormSqlite()
	default:
		global.RLB_ACTIVE_DBNAME = &global.RLB_CONFIG.Mysql.Dbname
		return server.GormMysql()
	}
}

func RegisterTables() {
	db := global.RLB_DB
	err := db.AutoMigrate(

	//TODO 初始化表

	)
	if err != nil {
		global.RLB_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	err = server.BizModel()

	if err != nil {
		global.RLB_LOG.Error("register biz_table failed", zap.Error(err))
		os.Exit(0)
	}
	global.RLB_LOG.Info("register table success")
}
