/*
 * @Author: Yang
 * @Date: 2025-04-05 09:35:03
 * @Description: 部门接口请求
 */
package request

import (
	common "github.com/kesilent/react-light-blog/dal/common/request"
)

type GetDeptListReq struct {
	common.PageInfo
	Path     string `json:"path" form:"path"`
	DeptName string `json:"deptName" form:"deptName"`
	Status   string `json:"status" form:"status"`
}
