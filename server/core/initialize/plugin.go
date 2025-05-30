package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/global"
)

func InstallPlugin(PrivateGroup *gin.RouterGroup, PublicRouter *gin.RouterGroup, engine *gin.Engine) {
	if global.RLB_DB == nil {
		global.RLB_LOG.Info("项目暂未初始化，无法安装插件，初始化后重启项目即可完成插件安装")
		return
	}
	bizPlugin(engine)
}
