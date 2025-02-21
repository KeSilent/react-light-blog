package global

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/config"
	"github.com/kesilent/react-light-blog/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	RLB_DB     *gorm.DB
	RLB_DBList map[string]*gorm.DB

	RLB_VP *viper.Viper

	RLB_CONFIG              config.Server //配置文件
	BlackCache              local_cache.Cache
	RLB_LOG                 *zap.Logger
	RLB_ACTIVE_DBNAME       *string
	RLB_Timer               timer.Timer = timer.NewTimerTask()
	RLB_Concurrency_Control             = &singleflight.Group{}
	RLB_ROUTERS             gin.RoutesInfo
)
