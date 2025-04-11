/*
 * @Author: Yang
 * @Date: 2025-03-19 19:09:08
 * @Description: 初始化数据
 */
package initializer

import (
	"github.com/kesilent/react-light-blog/dal/model"
	"github.com/kesilent/react-light-blog/dal/request"
	"github.com/kesilent/react-light-blog/services/system"
)

type InitData struct{}

func (i *InitData) Initialize() error {
	// 按顺序初始化：角色 -> 菜单 -> 用户 -> 关联关系
	if err := i.initAuthorities(); err != nil {
		return err
	}
	if err := i.initMenus(); err != nil {
		return err
	}
	if err := i.initUsers(); err != nil {
		return err
	}
	if err := i.initAuthorityMenus(); err != nil {
		return err
	}
	return i.initUserAuthorities()
}

func (i *InitData) initUsers() error {
	entities := []model.SysUser{
		{
			ID:       1,
			Username: "admin",
			Password: "123123",
			NickName: "超级管理员",
			Phone:    "18888888888",
			Email:    "admin@example.com",
			Avatar:   "https://os.alipayobjects.com/rmsportal/UXamdIxYSkXfoVo.jpg",
			Status:   true,
		},
	}

	for _, entity := range entities {
		if _, err := system.UserServiceApp.Register(entity); err != nil {
			return err
		}
	}
	return nil
}

func (i *InitData) initMenus() error {
	entities := []*model.SysBaseMenu{
		{ID: 1, Path: "dashboard", Name: "Dashboard", Component: "@/pages/dashboard/index", Sort: 1, Title: "仪表盘", Icon: "DashboardOutlined"},
		{ID: 2, Path: "system", Name: "System", Component: "@/layouts/RouteView", Sort: 2, Title: "系统管理", Icon: "SettingOutlined"},
		{ID: 3, ParentID: 2, Path: "/system/menu", Name: "Menu", Component: "./menu", Sort: 1, Title: "菜单管理", Icon: "MenuOutlined"},
		{ID: 4, ParentID: 2, Path: "/system/role", Name: "Role", Component: "./role", Sort: 2, Title: "角色管理", Icon: "TeamOutlined"},
		{ID: 5, ParentID: 2, Path: "/system/user", Name: "User", Component: "./user/manage", Sort: 3, Title: "用户管理", Icon: "UserOutlined"},
		{ID: 6, ParentID: 2, Path: "/system/dept", Name: "Dept", Component: "./dept", Sort: 4, Title: "部门管理", Icon: "AlibabaOutlined"},
	}
	return system.MenuServiceApp.AddBaseMenuList(entities)
}

func (i *InitData) initAuthorities() error {
	entities := []*model.SysRole{
		{ID: 888, RoleName: "超级管理员", DefaultRouter: "dashboard", ParentID: model.SnowflakeID(0)},
		{ID: 999, RoleName: "普通用户", DefaultRouter: "dashboard", ParentID: model.SnowflakeID(0)},
	}
	return system.RoleServiceApp.CreateRoleList(entities)
}

func (i *InitData) initAuthorityMenus() error {
	// 超级管理员拥有所有菜单权限
	adminMenus := []*model.SysRoleMenu{
		{SysRoleID: 888, SysBaseMenuID: 1},
		{SysRoleID: 888, SysBaseMenuID: 2},
		{SysRoleID: 888, SysBaseMenuID: 3},
		{SysRoleID: 888, SysBaseMenuID: 4},
		{SysRoleID: 888, SysBaseMenuID: 5},
	}
	// 普通用户只有仪表盘权限
	userMenus := []*model.SysRoleMenu{
		{SysRoleID: 999, SysBaseMenuID: 1},
	}

	entities := append(adminMenus, userMenus...)
	return system.RoleServiceApp.AddRoleMenus(request.RoleMenuReq{Rolemenus: entities, RoleID: 0})
}

func (i *InitData) initUserAuthorities() error {
	entities := []*model.SysUserRole{
		{SysUserID: 1, SysRoleID: 888},
		{SysUserID: 2, SysRoleID: 999},
	}
	return system.UserServiceApp.AddUserRole(entities)
}
