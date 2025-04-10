import { PageList, ResponseResult } from '@/models/system/common-model';
import { MenuModel } from '@/models/system/menu-model';
import { request } from '@umijs/max';

/**
 * @Author: Yang
 * @description: 通过关键词获取全部菜单
 * @return {*}
 */
export async function getMenuByKey(params: { keyWord: string }) {
  const result = await request<PageList<MenuModel[]>>('/api/menu/getMenuByKey', {
    method: 'GET',
    params: {
      ...params,
    },
  });

  return result.data || [];
}

/**
 * @Author: Yang
 * @description: 获取菜单列表
 * @return {*}
 */
export async function getMenuList(
  params: {
    current?: number;
    pageSize?: number;
  },
  options?: { [keyword: string]: any },
) {
  const result = await request<ResponseResult<PageList<MenuModel[]>>>(
    '/api/menu/getMenuListByPage',
    {
      method: 'GET',
      params: {
        ...params,
      },
      ...(options || {}),
    },
  );
  return result.data;
}

/**
 * @Author: Yang
 * @description: 删除菜单信息
 * @param {string} menuId
 * @return {*}
 */
export async function deleteMenu(menuId: string) {
  const result = await request<ResponseResult<string>>('/api/menu/deleteMenu', {
    method: 'DELETE',
    params: {
      id: menuId,
    },
  });
  return result.code;
}

/**
 * @Author: Yang
 * @description: 保存菜单信息
 * @param {MenuModel} params
 * @return {*}
 */
export async function saveMenu(params: MenuModel) {
  const result = await request<ResponseResult<string>>('/api/menu/saveBaseMenu', {
    method: 'POST',
    data: params,
  });
  return result.data;
}
