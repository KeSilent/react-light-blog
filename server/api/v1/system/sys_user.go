package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/models/common/response"
)

func (b *BaseApi) Login(c *gin.Context) {

	response.FailWithMessage("用户名不存在或者密码错误", c)
}
