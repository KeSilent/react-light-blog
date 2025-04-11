/*
 * @Author: Yang
 * @Date: 2025-04-03 22:25:53
 * @Description: 角色菜单绑定
 */
package request

import "github.com/kesilent/react-light-blog/dal/model"

type RoleMenuReq struct {
	Rolemenus []*model.SysRoleMenu `json:"rolemenus"`
	RoleID    model.SnowflakeID    `json:"roleID"`
}
