package system

import api "github.com/kesilent/react-light-blog/api/v1"

type RouterGroup struct {
	ApiRouter
	BaseRouter
}

var (
	baseApi = api.ApiGroupApp.SystemApiGroup.BaseApi
)
