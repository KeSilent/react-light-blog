package system

import "github.com/kesilent/react-light-blog/services"

type ApiGroup struct {
	BaseApi
}

var (
	userService = services.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService  = services.ServiceGroupApp.SystemServiceGroup.JwtService
)
