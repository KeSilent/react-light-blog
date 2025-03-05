package system

import "github.com/kesilent/react-light-blog/services"

type ApiGroup struct {
	BaseApi
	MenuApi
}

var (
	userService = services.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService  = services.ServiceGroupApp.SystemServiceGroup.JwtService
	menuService = services.ServiceGroupApp.SystemServiceGroup.MenuService
)
