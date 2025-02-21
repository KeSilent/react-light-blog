package v1

import "github.com/kesilent/react-light-blog/api/v1/system"

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}
