/*
 * @Author: Yang
 * @Date: 2025-03-19 19:09:08
 * @Description: 菜单操作类
 */
package system

import (
	"context"
	"errors"
	"sort"

	"github.com/kesilent/react-light-blog/dal/model"
	"github.com/kesilent/react-light-blog/dal/query"
	req "github.com/kesilent/react-light-blog/dal/request"
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
func (menuService *MenuService) GetRoleMenuList(roleUUID string) ([]model.SysBaseMenu, error) {
	// 获取权限信息，并预加载菜单

	q := query.Q.SysRole.WithContext(context.Background())

	if roleUUID != "" {
		q = q.Where(query.SysRole.UUID.Eq(roleUUID))
	}

	authority, err := q.
		Preload(field.NewRelation("Menus", "")).
		Preload(field.NewRelation("Menus.Children", "")).Debug().
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

/**
 * @Author: Yang
 * @description: 通过关键字查询菜单
 * @param {string} menuName
 * @return {[]model.SysBaseMenu, error}
 */
func (menuService *MenuService) GetMenuByKey(menuName string) ([]model.SysBaseMenu, error) {
	// 获取权限信息，并预加载菜单

	q := query.Q.SysBaseMenu.WithContext(context.Background())

	if menuName != "" {
		q = q.Where(query.SysBaseMenu.Name.Like("%" + menuName + "%"))
	}

	menus, err := q.Find()

	if err != nil {
		return nil, err
	}

	treeMenus := buildTree(menus, 0)

	// 对每个级别的菜单按 Sort 排序
	sortMenus(treeMenus)

	return treeMenus, nil
}

/**
 * @Author: Yang
 * @description:
 * @param {req.GetMenuListReq} info
 * @return {*}
 */
func (menuService *MenuService) GetMenuListByPage(info req.GetMenuListReq) (list []model.SysBaseMenu, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Current - 1)

	db := query.Q.SysBaseMenu.WithContext(context.Background())

	if info.Name != "" {
		db = db.Where(query.SysBaseMenu.Name.Like("%" + info.Name + "%"))
	}

	total, err = db.Count()
	if err != nil {
		return
	}

	userList, err := db.Limit(limit).Offset(offset).Find()

	treeMenus := buildTree(userList, 0)

	// 对每个级别的菜单按 Sort 排序
	sortMenus(treeMenus)

	return treeMenus, int64(len(treeMenus)), err
}

// 构建树形结构
func buildTree(menus []*model.SysBaseMenu, parentID int64) []model.SysBaseMenu {
	var tree []model.SysBaseMenu

	for _, menu := range menus {
		if menu.ParentID == parentID {
			// 复制当前菜单节点
			node := *menu
			// 递归构建子菜单
			node.Children = buildTree(menus, menu.ID)
			tree = append(tree, node)
		}
	}

	return tree
}
