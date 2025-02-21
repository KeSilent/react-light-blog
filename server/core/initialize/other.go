package initialize

import (
	"github.com/kesilent/react-light-blog/global"
	"github.com/kesilent/react-light-blog/utils"
	"github.com/songzhibin97/gkit/cache/local_cache"
)

func OtherInit() {
	//JWT时间
	dr, err := utils.ParseDuration(global.RLB_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.RLB_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)

}
