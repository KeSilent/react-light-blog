/*
 * @Author: Yang
 * @Date: 2025-04-05 09:35:03
 * @Description: 菜单接口请求
 */
package request

import (
	common "github.com/kesilent/react-light-blog/dal/common/request"
	"github.com/kesilent/react-light-blog/dal/model"
)

type GetMenuListReq struct {
	common.PageInfo
	model.SysBaseMenu
}
