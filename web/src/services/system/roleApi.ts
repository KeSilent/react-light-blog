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
