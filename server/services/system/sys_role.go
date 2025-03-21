/*
 * @Author: Yang
 * @Date: 2025-03-19 19:09:08
 * @Description: 角色操作类
 */
package system

import (
	"context"

	"github.com/kesilent/react-light-blog/dal/model"
	"github.com/kesilent/react-light-blog/dal/query"
	systemReq "github.com/kesilent/react-light-blog/dal/request"
	"gorm.io/gen"
)

type RoleService struct{}

var RoleServiceApp = new(RoleService)

/**
 * @author: JackYang
 * @function: CreateRoleList
 * @description: 角色批量创建
 * @param: authorities []*model.SysRole
 * @return: error
 */
func (role *RoleService) CreateRoleList(authorities []*model.SysRole) error {
	q := query.SysRole.WithContext(context.Background())
	return q.Create(authorities...)
}

/**
 * @author: JackYang
 * @function: AddRoleMenus
 * @description: 角色菜单绑定
 * @param: authorityMenus []*model.SysRoleMenu
 * @return: error
 */
func (role *RoleService) AddRoleMenus(authorityMenus []*model.SysRoleMenu) error {
	q := query.SysRoleMenu
	if _, err := role.DeleteRoleMenu(authorityMenus[0].SysRoleRoleID); err != nil {
		return err
	}
	if authorityMenus[0].SysBaseMenuID != "" {
		return q.WithContext(context.Background()).Create(authorityMenus...)
	}
	return nil
}

func (role *RoleService) DeleteRoleMenu(roleUUID string) (gen.ResultInfo, error) {
	q := query.SysRoleMenu
	return q.WithContext(context.Background()).Where(q.SysRoleRoleID.Eq(roleUUID)).Delete()
}

/**
 * @author: JackYang
 * @function: GetRoleMenus
 * @description: 获取角色菜单
 * @param: roleId int64
 * @return: []*model.SysRoleMenu, error
 */
func (role *RoleService) GetRoleMenus(roleUUID string) ([]*model.SysRoleMenu, error) {
	q := query.SysRoleMenu
	return q.WithContext(context.Background()).Where(q.SysRoleRoleID.Eq(roleUUID)).Find()
}

// @author: JackYang
// @function: GetRoleList
// @description: 分页获取数据
// @param: info request.GetRoleListReq
// @return: err error, list interface{}, total int64
func (role *RoleService) GetRoleList(info systemReq.GetRoleListReq) (list []*model.SysRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Current - 1)

	db := query.Q.SysRole.WithContext(context.Background())

	if info.Name != "" {
		db = db.Where(query.SysRole.RoleName.Like("%" + info.Name + "%"))
	}

	total, err = db.Count()
	if err != nil {
		return
	}

	userList, err := db.Limit(limit).Offset(offset).Find()

	return userList, total, err
}
