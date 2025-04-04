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
			UUID:     "815b6a0e-756e-4aa4-ac3c-e271518a5f93",
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
		{ID: 1, UUID: "f1344409-c286-492f-be6c-d6798d5d3e6f", Path: "dashboard", Name: "Dashboard", Component: "@/pages/dashboard/index", Sort: 1, Title: "仪表盘", Icon: "DashboardOutlined"},
		{ID: 2, UUID: "ec350c05-78d8-4b2a-b588-6d71ed0b9960", Path: "system", Name: "System", Component: "@/layouts/RouteView", Sort: 2, Title: "系统管理", Icon: "SettingOutlined"},
		{ID: 3, UUID: "e84d2d69-5300-45d0-9ad9-811898f8440e", ParentID: 2, Path: "menu", Name: "Menu", Component: "@/pages/system/menu/index", Sort: 1, Title: "菜单管理", Icon: "MenuOutlined"},
		{ID: 4, UUID: "b51cf3a4-73ca-4981-9f28-dc2e7521dcba", ParentID: 2, Path: "role", Name: "Role", Component: "@/pages/system/role/index", Sort: 2, Title: "角色管理", Icon: "TeamOutlined"},
		{ID: 5, UUID: "dc6f1752-25d3-41aa-8764-8de484bee73c", ParentID: 2, Path: "user", Name: "User", Component: "@/pages/system/user/index", Sort: 3, Title: "用户管理", Icon: "UserOutlined"},
	}
	return system.MenuServiceApp.AddBaseMenuList(entities)
}

func (i *InitData) initAuthorities() error {
	entities := []*model.SysRole{
		{ID: 888, UUID: "9f60b4b9-ffc1-4e83-aa80-9d48ff291928", RoleName: "超级管理员", DefaultRouter: "dashboard"},
		{ID: 999, UUID: "ad43c17d-69a7-4b05-aa09-f5577892c684", RoleName: "普通用户", ParentID: 888, DefaultRouter: "dashboard"},
	}
	return system.RoleServiceApp.CreateRoleList(entities)
}

func (i *InitData) initAuthorityMenus() error {
	// 超级管理员拥有所有菜单权限
	adminMenus := []*model.SysRoleMenu{
		{SysRoleUUID: "9f60b4b9-ffc1-4e83-aa80-9d48ff291928", SysBaseMenuUUID: "b51cf3a4-73ca-4981-9f28-dc2e7521dcba"},
		{SysRoleUUID: "9f60b4b9-ffc1-4e83-aa80-9d48ff291928", SysBaseMenuUUID: "dc6f1752-25d3-41aa-8764-8de484bee73c"},
		{SysRoleUUID: "9f60b4b9-ffc1-4e83-aa80-9d48ff291928", SysBaseMenuUUID: "e84d2d69-5300-45d0-9ad9-811898f8440e"},
		{SysRoleUUID: "9f60b4b9-ffc1-4e83-aa80-9d48ff291928", SysBaseMenuUUID: "ec350c05-78d8-4b2a-b588-6d71ed0b9960"},
		{SysRoleUUID: "9f60b4b9-ffc1-4e83-aa80-9d48ff291928", SysBaseMenuUUID: "f1344409-c286-492f-be6c-d6798d5d3e6f"},
	}
	// 普通用户只有仪表盘权限
	userMenus := []*model.SysRoleMenu{
		{SysRoleUUID: "c311fdcc-65ae-4140-96a1-846dd54ba9ca", SysBaseMenuUUID: "b51cf3a4-73ca-4981-9f28-dc2e7521dcba"},
	}

	entities := append(adminMenus, userMenus...)
	return system.RoleServiceApp.AddRoleMenus(request.RoleMenuReq{Rolemenus: entities, RoleUUID: ""})
}

func (i *InitData) initUserAuthorities() error {
	entities := []*model.SysUserRole{
		{SysUserUUID: "815b6a0e-756e-4aa4-ac3c-e271518a5f93", SysRoleUUID: "9f60b4b9-ffc1-4e83-aa80-9d48ff291928"},
		{SysUserUUID: "c311fdcc-65ae-4140-96a1-846dd54ba9ca", SysRoleUUID: "ad43c17d-69a7-4b05-aa09-f5577892c684"},
	}
	return system.UserServiceApp.AddUserRole(entities)
}
