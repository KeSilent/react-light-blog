package system

import "github.com/kesilent/react-light-blog/services"

type ApiGroup struct {
	BaseApi
	MenuApi
	RoleApi
	DeptApi
}

var (
	userService = services.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService  = services.ServiceGroupApp.SystemServiceGroup.JwtService
	menuService = services.ServiceGroupApp.SystemServiceGroup.MenuService
	roleService = services.ServiceGroupApp.SystemServiceGroup.RoleService
	deptService = services.ServiceGroupApp.SystemServiceGroup.DeptService
)
