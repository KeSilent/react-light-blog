package core

import (
	"fmt"

	"github.com/kesilent/react-light-blog/core/initialize"
	"github.com/kesilent/react-light-blog/global"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	if global.RLB_CONFIG.System.UseMultipoint || global.RLB_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
		initialize.RedisList()
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.RLB_CONFIG.System.Addr)
	s := initServer(address, Router)

	global.RLB_LOG.Info("server run success on ", zap.String("address", address))

	global.RLB_LOG.Error(s.ListenAndServe().Error())
}
