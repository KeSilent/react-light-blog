package system

import (
	"context"
	"errors"

	"github.com/kesilent/react-light-blog/dal/model"
	"github.com/kesilent/react-light-blog/dal/query"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type MenuService struct{}

var MenuServiceApp = new(MenuService)

// @author: JackYang
// @function: SysBaseMenu
// @description: 用户角色默认路由检查
// @param: 参数类型
// @return: 返回类型
func (menuService *MenuService) UserAuthorityDefaultRouter(authority *model.SysAuthority) {
	var menuIds []int64

	db := query.Q.SysAuthorityMenu.WithContext(context.Background())
	err := db.Where(query.SysAuthorityMenu.SysAuthorityAuthorityID.Eq(authority.ID)).Pluck(query.SysAuthorityMenu.SysBaseMenuID, &menuIds)

	if err != nil {
		return
	}
	menuDB := query.Q.SysBaseMenu.WithContext(context.Background())
	_, err = menuDB.Where(query.SysBaseMenu.ID.In(menuIds...), query.SysBaseMenu.Name.Eq(authority.DefaultRouter)).First()

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

// @author: JackYang
// @function: GetAuthorMenuList
// @description: 获取角色菜单
// @param: authorId int64
// @return: 返回类型
func (menuService *MenuService) GetAuthorMenuList(authorId int64) ([]model.SysBaseMenu, error) {
	// 获取权限信息，并预加载菜单
	authority, err := query.Q.SysAuthority.WithContext(context.Background()).
		Where(query.SysAuthority.ID.Eq(authorId)).
		Preload(field.NewRelation("Menus", "")).
		Preload(field.NewRelation("Menus.Children", "")).
		First()

	if err != nil {
		return nil, err
	}

	return authority.Menus, nil
}
