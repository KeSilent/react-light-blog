package initialize

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/global"
	"github.com/kesilent/react-light-blog/middleware"
	router "github.com/kesilent/react-light-blog/routes"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

//总路由

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	systemRouter := router.RouterGroupApp.System
	userRouter := router.RouterGroupApp.User
	roleRouter := router.RouterGroupApp.Role
	menuRouter := router.RouterGroupApp.Menu

	PublicGroup := Router.Group(global.RLB_CONFIG.System.RouterPrefix)
	PrivateGroup := Router.Group(global.RLB_CONFIG.System.RouterPrefix)

	// Router.StaticFS(global.RLB_CONFIG.Local.StorePath, justFilesFilesystem{http.Dir(global.RLB_CONFIG.Local.StorePath)}) // Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")

	PrivateGroup.Use(middleware.JWTAuth())

	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		userRouter.InitUserRouter(PublicGroup)
		roleRouter.InitRoleRouter(PublicGroup)
		menuRouter.InitMenuRouter(PublicGroup)
	}

	global.RLB_ROUTERS = Router.Routes()
	global.RLB_LOG.Info("router register success")
	return Router
}
