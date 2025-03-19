package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/middleware"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("user")
	authorityRouter := Router.Group("user").Use(middleware.OperationRecord())
	{
		baseRouter.POST("register", baseApi.Register)
	}
	{
		authorityRouter.GET("getUserList", baseApi.GetUserList)
		authorityRouter.POST("changePassword", baseApi.ChangePassword)
		authorityRouter.POST("updateUser", baseApi.UpdateUser)
	}
	return baseRouter
}
