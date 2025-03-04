package initializer

import (
	"github.com/kesilent/react-light-blog/dal/model"
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
			ID:          1,
			Username:    "admin",
			Password:    "123123",
			NickName:    "超级管理员",
			AuthorityID: 888,
			Phone:       "18888888888",
			Email:       "admin@example.com",
			HeaderImg:   "https://os.alipayobjects.com/rmsportal/UXamdIxYSkXfoVo.jpg",
			Status:      true,
		},
		{
			ID:          2,
			Username:    "test",
			Password:    "123123",
			NickName:    "测试用户",
			AuthorityID: 999,
			Phone:       "13999999999",
			HeaderImg:   "https://os.alipayobjects.com/rmsportal/UXamdIxYSkXfoVo.jpg",
			Email:       "test@example.com",
			Status:      true,
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
		{ID: 3, ParentID: 2, Path: "menu", Name: "Menu", Component: "@/pages/system/menu/index", Sort: 1, Title: "菜单管理", Icon: "MenuOutlined"},
		{ID: 4, ParentID: 2, Path: "role", Name: "Role", Component: "@/pages/system/role/index", Sort: 2, Title: "角色管理", Icon: "TeamOutlined"},
		{ID: 5, ParentID: 2, Path: "user", Name: "User", Component: "@/pages/system/user/index", Sort: 3, Title: "用户管理", Icon: "UserOutlined"},
	}
	return system.MenuServiceApp.AddBaseMenuList(entities)
}

func (i *InitData) initAuthorities() error {
	entities := []*model.SysAuthority{
		{ID: 888, AuthorityName: "超级管理员", DefaultRouter: "dashboard"},
		{ID: 999, AuthorityName: "普通用户", ParentID: 888, DefaultRouter: "dashboard"},
	}
	return system.AuthorityServiceApp.CreateAuthorityList(entities)
}

func (i *InitData) initAuthorityMenus() error {
	// 超级管理员拥有所有菜单权限
	adminMenus := []*model.SysAuthorityMenu{
		{SysAuthorityID: 888, SysBaseMenuID: 1},
		{SysAuthorityID: 888, SysBaseMenuID: 2},
		{SysAuthorityID: 888, SysBaseMenuID: 3},
		{SysAuthorityID: 888, SysBaseMenuID: 4},
		{SysAuthorityID: 888, SysBaseMenuID: 5},
	}
	// 普通用户只有仪表盘权限
	userMenus := []*model.SysAuthorityMenu{
		{SysAuthorityID: 999, SysBaseMenuID: 1},
	}

	entities := append(adminMenus, userMenus...)
	return system.AuthorityServiceApp.AddAuthorityMenus(entities)
}

func (i *InitData) initUserAuthorities() error {
	entities := []*model.SysUserAuthority{
		{SysUserID: 1, SysAuthorityID: 888},
		{SysUserID: 2, SysAuthorityID: 999},
	}
	return system.UserServiceApp.AddUserAuthorities(entities)
}
