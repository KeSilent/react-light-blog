package system

import api "github.com/kesilent/react-light-blog/api/v1"

type RouterGroup struct {
	ApiRouter
	BaseRouter
	UserRouter
	MenuRouter
	AutoCodeRouter
}

var (
	baseApi             = api.ApiGroupApp.SystemApiGroup.BaseApi
	menuApi             = api.ApiGroupApp.SystemApiGroup.MenuApi
	roleApi             = api.ApiGroupApp.SystemApiGroup.RoleApi
	deptApi             = api.ApiGroupApp.SystemApiGroup.DeptApi
	codeBuilderFieldApi = api.ApiGroupApp.SystemApiGroup.CodeBuilderFieldApi
	AutoCodeApi         = api.ApiGroupApp.SystemApiGroup.AutoCodeApi
)
