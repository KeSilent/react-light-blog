package request

import (
	common "github.com/kesilent/react-light-blog/dal/common/request"
)

type GetRoleListReq struct {
	common.PageInfo
	RoleName string `json:"roleName" form:"roleName"`
}
