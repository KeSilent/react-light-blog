package system

import (
	"errors"

	"github.com/kesilent/react-light-blog/global"
	"github.com/kesilent/react-light-blog/models/system"
	"gorm.io/gorm"
)

type MenuService struct{}

var MenuServiceApp = new(MenuService)

// @author: JackYang
// @function: SysMenu
// @description: 用户角色默认路由检查
// @param: 参数类型
// @return: 返回类型
func (menuService *MenuService) UserAuthorityDefaultRouter(user *system.SysUser) {
	var menuIds []string
	err := global.RLB_DB.Model(&system.SysAuthorityMenu{}).Where("sys_authority_authority_id = ?", user.AuthorityId).Pluck("sys_base_menu_id", &menuIds).Error
	if err != nil {
		return
	}
	var am system.SysBaseMenu
	err = global.RLB_DB.First(&am, "name = ? and id in (?)", user.Authority.DefaultRouter, menuIds).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Authority.DefaultRouter = "404"
	}
}
