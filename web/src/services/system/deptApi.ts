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
  const result = await request<ResponseResult<PageList<DeptModel[]>>>('/api/dept/getList', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
  return result.data;
}

/**
 * @Author: Yang
 * @description: 获取下拉菜单部门列表
 * @param {object} params
 * @return {*}
 */
export async function getListByTreeSelect(params: { keyWord: string }) {
  const result = await request<PageList<DeptModel[]>>('/api/dept/getListByTreeSelect', {
    method: 'GET',
    params: {
      ...params,
    },
  });

  return result.data || [];
}

/**
 * @Author: Yang
 * @description: 删除部门信息
 * @param {string} deptId
 * @return {*}
 */
export async function deleteDept(deptId: string) {
  const result = await request<ResponseResult<string>>('/api/dept/deleteDept', {
    method: 'DELETE',
    params: {
      id: deptId,
    },
  });
  return result.code;
}

/**
 * @Author: Yang
 * @description: 保存部门信息
 * @param {DeptModel} params
 * @return {*}
 */
export async function saveDept(params: DeptModel) {
  const result = await request<ResponseResult<string>>('/api/dept/saveDept', {
    method: 'POST',
    data: params,
  });
  return { data: result.code };
}
