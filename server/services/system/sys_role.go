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
 * @function: SaveRole
 * @description: 保存角色
 * @param: model *model.SysRole
 * @return: error
 */
func (role *RoleService) SaveRole(model *model.SysRole) error {
	return query.SysRole.WithContext(context.Background()).Save(model)
}

/**
 * @author: JackYang
 * @function: DeleteRole
 * @description: 删除角色
 * @param: id int64
 * @return: gen.ResultInfo, error
 */
func (role *RoleService) DeleteRole(id int64) (gen.ResultInfo, error) {
	q := query.SysRole.WithContext(context.Background())
	return q.Where(query.SysRole.ID.Eq(model.SnowflakeID(id))).Delete()
}

/**
 * @author: JackYang
 * @function: AddRoleMenus
 * @description: 角色菜单绑定
 * @param: authorityMenus []*model.SysRoleMenu
 * @return: error
 */
func (role *RoleService) AddRoleMenus(authorityMenus systemReq.RoleMenuReq) error {
	q := query.SysRoleMenu
	if _, err := role.DeleteRoleMenu(int64(authorityMenus.RoleID)); err != nil {
		return err
	}
	if len(authorityMenus.Rolemenus) > 0 {
		return q.WithContext(context.Background()).Create(authorityMenus.Rolemenus...)
	}
	return nil
}

func (role *RoleService) DeleteRoleMenu(roleID int64) (gen.ResultInfo, error) {
	q := query.SysRoleMenu
	return q.WithContext(context.Background()).Where(q.SysRoleID.Eq(model.SnowflakeID(roleID))).Delete()
}

/**
 * @author: JackYang
 * @function: GetRoleMenus
 * @description: 获取角色菜单
 * @param: roleId int64
 * @return: []*model.SysRoleMenu, error
 */
func (role *RoleService) GetRoleMenus(roleID int64) ([]*model.SysRoleMenu, error) {
	q := query.SysRoleMenu
	return q.WithContext(context.Background()).Where(q.SysRoleID.Eq(model.SnowflakeID(roleID))).Find()
}

/**
 * @author: JackYang
 * @function: GetRoleMenus
 * @description: 分页获取数据
 * @param: roleId int64
 * @return: []*model.SysRoleMenu, error
 */
func (role *RoleService) GetRoleList(info systemReq.GetRoleListReq) (list []*model.SysRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Current - 1)

	db := query.Q.SysRole.WithContext(context.Background())

	if info.RoleName != "" {
		db = db.Where(query.SysRole.RoleName.Like("%" + info.RoleName + "%"))
	}

	total, err = db.Count()
	if err != nil {
		return
	}

	userList, err := db.Limit(limit).Offset(offset).Find()

	return userList, total, err
}
