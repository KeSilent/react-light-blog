package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/middleware"
)

type RoleRouter struct{}

func (s *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	authorityRouter := Router.Group("role").Use(middleware.OperationRecord())
	{
		authorityRouter.GET("getRoleList", roleApi.GetRoleList)
	}
	return authorityRouter
}
