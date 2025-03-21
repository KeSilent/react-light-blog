package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/dal/common/response"
	"github.com/kesilent/react-light-blog/utils"
)

type MenuApi struct{}

func (m *MenuApi) Menus(c *gin.Context) {
	user, _ := userService.GetUserInfo(utils.GetUserUuid(c))
	if len(user.Role) == 0 {
		response.FailWithMessage("未设置任何角色", c)
	}

	//获取第一个角色
	menus, err := menuService.GetRoleMenuList(user.Role[0].UUID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(menus, c)
		return
	}
}

/***
 * @Author: kesilent
 * @Description: 获取菜单列表
 **/
func (m *MenuApi) GetMenuList(c *gin.Context) {
	keyWord := c.Query("keyWord")
	//获取列表
	menus, err := menuService.GetMenuList(keyWord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(menus, c)
		return
	}
}
