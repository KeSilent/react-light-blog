package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/middleware"
)

type CodeBuilderFieldRouter struct{}

func (d *CodeBuilderFieldRouter) InitCodeBuilderFieldRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	deptRouter := Router.Group("code").Use(middleware.OperationRecord())
	{
		deptRouter.GET("getListByPage", codeBuilderFieldApi.GetListByPage)
	}
	return deptRouter
}
