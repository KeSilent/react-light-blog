/*
 * @Author: Yang
 * @Date: 2025-04-05 09:35:03
 * @Description: 菜单接口请求
 */
package request

import (
	common "github.com/kesilent/react-light-blog/dal/common/request"
)

type GetMenuListReq struct {
	common.PageInfo
	Path      string `json:"path" form:"path"`
	Name      string `json:"name" form:"name"`
	Component string `json:"component" form:"component"` // 对应前端文件路径
	Title     string `json:"title" form:"title"`         // 菜单名
}
