import { PageList, ResponseResult } from '@/models/system/common-model';
import { RoleMenuModel, RoleModel } from '@/models/system/role-model';
import { request } from '@umijs/max';

/**
 * 获取角色列表
 * @param params
 * @param options
 * @returns
 */
export async function getRoleList(
  params: {
    current?: number;
    pageSize?: number;
  },
  options?: { [keyword: string]: any },
) {
  const result = await request<ResponseResult<PageList<RoleModel[]>>>('/api/role/getRoleList', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });

  return result.data;
}

export async function getAllRoleKeyValueList() {
  const result = await request<ResponseResult<PageList<RoleModel[]>>>('/api/role/getRoleList', {
    method: 'GET',
    params: {
      current: 0,
      pageSize: 999,
    },
  });
  if (result.data) {
    if (result.data.data) {
      return result.data.data.map((item) => {
        return {
          label: item.roleName,
          value: item.id,
        };
      });
    }
  }
  return [];
}

/**
 * 获取角色菜单
 * @param roleId
 * @returns
 */
export async function getRoleMenus(roleId: string) {
  const result = await request<ResponseResult<RoleMenuModel[]>>('/api/role/getRoleMenus', {
    method: 'GET',
    params: {
      roleId: roleId,
    },
  });
  return result.data;
}

/**
 * 添加角色菜单
 * @param params
 * @returns
 */
export async function addRoleMenu(params: RoleMenuModel) {
  const result = await request<ResponseResult<RoleModel>>('/api/role/addRoleMenu', {
    method: 'POST',
    data: params,
  });
  return { data: result.code === 0 };
}

/**
 * 保存角色信息
 * @param params 角色信息
 * @returns
 */
export async function saveRole(params: RoleModel) {
  const result = await request<ResponseResult<RoleModel>>('/api/role/saveRole', {
    method: 'POST',
    data: params,
  });
  return { data: result.code === 0 };
}
/**
 * 删除角色信息
 * @param params 角色信息
 * @returns
 */
export async function deleteRole(roleUUId: string) {
  const result = await request<ResponseResult>('/api/role/deleteRole', {
    method: 'DELETE',
    params: { id: roleUUId },
  });
  return result.code;
}
