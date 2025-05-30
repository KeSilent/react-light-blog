package request

import (
	common "{{.Module}}/dal/common/request"
)

type Get{{.StructName}}ListReq struct {
	common.PageInfo
}
