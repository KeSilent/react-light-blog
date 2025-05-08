import { CodeBuilderFieldModel } from '@/models/system/code-builder-fields-model';
import { PageList, ResponseResult } from '@/models/system/common-model';
import { request } from '@umijs/max';
/**
 * @Author: Yang
 * @description: 获取代码字段列表
 * @return {*}
 */
export async function getCodeFieldList(
  params: {
    current?: number;
    pageSize?: number;
  },
  options?: { [keyword: string]: any },
) {
  const result = await request<ResponseResult<PageList<CodeBuilderFieldModel[]>>>(
    '/api/code/getListByPage',
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
 * @description: 删除字段
 * @param {string} fieldId
 * @return {*}
 */
export async function deleteField(fieldId: string) {
  const result = await request<ResponseResult<string>>('/api/menu/deleteMenu', {
    method: 'DELETE',
    params: {
      id: fieldId,
    },
  });
  return result.code;
}
