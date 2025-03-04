import { request } from '@umijs/max';

export async function getDynamicRoutes() {
  return request('/api/menus', {
    method: 'GET',
  });
}