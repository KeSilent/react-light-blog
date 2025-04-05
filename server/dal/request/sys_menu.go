/*
 * @Author: Yang
 * @Date: 2025-04-05 09:35:03
 * @Description: 菜单接口请求
 */
package request

import common "github.com/kesilent/react-light-blog/dal/common/request"

type GetMenuListReq struct {
	common.PageInfo
	Name string `json:"name" form:"name"`
}
