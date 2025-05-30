package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/middleware"
)

type AutoCodeRouter struct{}

func (d *AutoCodeRouter) InitAutoCodeRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	deptRouter := Router.Group("autoCode").Use(middleware.OperationRecord())
	{
		deptRouter.GET("getAllTableName", AutoCodeApi.GetAllTableName)
		deptRouter.GET("getFieldsByTableName", AutoCodeApi.GetFieldsByTableName)
	}
	{
		deptRouter.POST("createTemp", AutoCodeApi.CreateTemp)
	}
	return deptRouter
}
