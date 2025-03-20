import { PageList, ResponseResult } from "@/models/system/common-model";
import { MenuModel } from "@/models/system/menu-model";
import { request } from "@umijs/max";

export async function getMenuList(
  params: {
    keyword: string
  },
) {
  const result = await request<PageList<MenuModel[]>>('/api/menu/list', {
    method: 'GET',
    params: {
      ...params,
    },
  });

  return result.data || [];
}