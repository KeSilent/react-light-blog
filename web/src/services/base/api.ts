
import { MenuModel } from '@/models/system/menu-model';
import { request } from '@umijs/max';


/** 登录接口 POST /api/login/account */
export async function login(body: API.LoginParams, options?: { [key: string]: any }) {
  return request<API.LoginResult>('/api/base/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/**
 * 获取权限菜单
 * @returns 
 */
export async function getDynamicMenus() {
  return request<API.PageResponse<MenuModel[]>>('/api/authority/menus', {
    method: 'GET',
  });
}
