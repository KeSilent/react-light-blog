package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/utils/plugin"
)

func PluginInit(group *gin.Engine, plugins ...plugin.Plugin) {
	for i := 0; i < len(plugins); i++ {
		plugins[i].Register(group)
	}
}
func bizPlugin(engine *gin.Engine) {
}
