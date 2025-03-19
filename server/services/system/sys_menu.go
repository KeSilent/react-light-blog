package system

import (
	"context"
	"errors"
	"sort"

	"github.com/kesilent/react-light-blog/dal/model"
	"github.com/kesilent/react-light-blog/dal/query"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type MenuService struct{}

var MenuServiceApp = new(MenuService)

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
// @function: GetRoleMenuList
// @description: 获取角色菜单
// @param: authorId int64
// @return: 返回类型
func (menuService *MenuService) GetRoleMenuList(authorId int64) ([]model.SysBaseMenu, error) {
	// 获取权限信息，并预加载菜单
	authority, err := query.Q.SysRole.WithContext(context.Background()).
		Where(query.SysRole.ID.Eq(authorId)).
		Preload(field.NewRelation("Menus", "")).
		Preload(field.NewRelation("Menus.Children", "")).
		First()

	if err != nil {
		return nil, err
	}
	var menus = authority.Menus
	// 构建树形结构
	menuMap := make(map[int64]*model.SysBaseMenu)
	var rootMenus []model.SysBaseMenu

	// 首先构建一个 ID 到菜单的映射
	for _, menu := range menus {
		menuCopy := menu
		menuMap[menu.ID] = &menuCopy
	}

	// 构建树形结构
	for _, menu := range menus {
		if menu.ParentID == 0 {
			// 这是根菜单
			rootMenus = append(rootMenus, menu)
		} else {
			// 找到父菜单并添加到其子菜单中
			if parent, ok := menuMap[menu.ParentID]; ok {
				parent.Children = append(parent.Children, menu)
			}
		}
	}

	// 对每个级别的菜单按 Sort 排序
	sortMenus(rootMenus)

	return rootMenus, nil
}

// 递归排序菜单
func sortMenus(menus []model.SysBaseMenu) {
	// 按 Sort 字段排序
	sort.Slice(menus, func(i, j int) bool {
		return menus[i].Sort < menus[j].Sort
	})

	// 递归排序子菜单
	for i := range menus {
		if len(menus[i].Children) > 0 {
			sortMenus(menus[i].Children)
		}
	}
}
