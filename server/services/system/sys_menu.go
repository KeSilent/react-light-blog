package system

import (
	"context"
	"errors"

	"github.com/kesilent/react-light-blog/dal/model"
	"github.com/kesilent/react-light-blog/dal/query"
	"gorm.io/gorm"
)

type MenuService struct{}

var MenuServiceApp = new(MenuService)

// @author: JackYang
// @function: SysBaseMenu
// @description: 用户角色默认路由检查
// @param: 参数类型
// @return: 返回类型
func (menuService *MenuService) UserAuthorityDefaultRouter(user *model.SysUser) {
	var menuIds []int64

	db := query.Q.SysAuthorityMenu.WithContext(context.Background())
	err := db.Where(query.SysAuthorityMenu.SysAuthorityID.Eq(user.AuthorityID)).Pluck(query.SysAuthorityMenu.SysBaseMenuID, &menuIds)

	if err != nil {
		return
	}
	menuDB := query.Q.SysBaseMenu.WithContext(context.Background())
	_, err = menuDB.Where(query.SysBaseMenu.ID.In(menuIds...), query.SysBaseMenu.Name.Eq(user.Authority.DefaultRouter)).First()

	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Authority.DefaultRouter = "404"
	}
}

//@author: JackYang
//@function: AddBaseMenu
//@description: 增加菜单
//@param: menu model.SysBaseMenu
//@return: error

func (menuService *MenuService) AddBaseMenu(menu model.SysBaseMenu) error {
	db := query.Q.SysBaseMenu.WithContext(context.Background())
	_, err := db.Where(query.SysBaseMenu.Name.Eq(menu.Name)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return db.Create(&menu)
}

func (menuService *MenuService) AddBaseMenuList(menus []*model.SysBaseMenu) error {
	db := query.Q.SysBaseMenu.WithContext(context.Background())
	return db.Create(menus...)
}
