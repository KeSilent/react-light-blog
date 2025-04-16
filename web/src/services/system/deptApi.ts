import { PageList, ResponseResult } from '@/models/system/common-model';
import { DeptModel } from '@/models/system/dept-model';
import { request } from '@umijs/max';

/**
 * @Author: Yang
 * @description: 获取部门列表
 * @return {*}
 */
export async function getDeptList(
  params: {
    current?: number;
    pageSize?: number;
  },
  options?: { [keyword: string]: any },
) {
  const result = await request<ResponseResult<PageList<DeptModel[]>>>(
    '/api/dept/getList',
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
 * @description: 删除部门信息
 * @param {string} deptId
 * @return {*}
 */
export async function deleteDept(deptId: string) {
  const result = await request<ResponseResult<string>>('/api/menu/deleteMenu', {
    method: 'DELETE',
    params: {
      id: deptId,
    },
  });
  return result.code;
}


export async function saveDept(params: DeptModel) {
  const result = await request<ResponseResult<string>>('/api/dept/saveDept', {
    method: 'POST',
    data: params,
  });
  return result.code;
}