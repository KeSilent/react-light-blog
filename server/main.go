package main

import (
	"github.com/kesilent/react-light-blog/core"
	"github.com/kesilent/react-light-blog/core/initialize"
	"github.com/kesilent/react-light-blog/dal/query"
	"github.com/kesilent/react-light-blog/global"
	"go.uber.org/zap"
)

func main() {
	global.RLB_VP = core.Viper() //初始化配置文件读取
	initialize.OtherInit()
	global.RLB_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.RLB_LOG)

	global.RLB_DB = initialize.Gorm() // gorm连接数据库

	initialize.Timer() //定时任务

	//TODO 生成代码，发布时候注释掉
	// core.GenStructs()

	if global.RLB_DB != nil {
		query.SetDefault(global.RLB_DB)
		// initialize.RegisterTables() // 初始化表

		// 程序结束前关闭数据库链接
		// db, _ := global.RLB_DB.DB()
		// defer db.Close()
	}
	core.RunWindowsServer()
}
