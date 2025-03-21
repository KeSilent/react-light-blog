import { PageList } from "@/models/system/common-model";
import { MenuModel } from "@/models/system/menu-model";
import { request } from "@umijs/max";

export async function getMenuList(
  params: {
    keyWord: string
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