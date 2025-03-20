import { PageList, ResponseResult } from "@/models/system/common-model";
import { RoleModel } from "@/models/system/role-model";
import { request } from "@umijs/max";

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

  return result.data
}