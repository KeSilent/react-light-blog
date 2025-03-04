package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/dal/common/response"
	"github.com/kesilent/react-light-blog/services"
)

func (i *BaseApi) InitData(c *gin.Context) {
	err := services.ServiceGroupApp.InitializerGroup.Initialize()

	if err != nil {
		response.FailWithMessage("初始化失败", c)
	} else {
		response.OkWithMessage("初始化成功", c)
	}

}
