package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/dal/common/response"
	"github.com/kesilent/react-light-blog/utils"
)

type MenuApi struct{}

func (m *MenuApi) Menus(c *gin.Context) {
	user, _ := userService.GetUserInfo(utils.GetUserID(c))
	if len(user.Role) == 0 {
		response.FailWithMessage("未设置任何角色", c)
	}

	//获取第一个角色
	menus, err := menuService.GetAuthorMenuList(user.Role[0].ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(menus, c)
		return
	}
}
