import { CodeFieldModel } from '@/models/system/code-field-model';
import {  ResponseResult } from '@/models/system/common-model';
import { request } from '@umijs/max';
/**
 * @Author: Yang
 * @description: 获取所有表
 * @return {*}
 */
export async function getAllTableName() {
  const result = await request<ResponseResult<string[]>>('/api/autoCode/getAllTableName', {
    method: 'GET',
  });

  return result.data || [];
}

/**
 * @Author: Yang
 * @description: 获取表中所有的字段
 * @param {string} tableName
 * @return {*}
 */
export async function getFieldsByTableName(tableName: string) {
  const result = await request<ResponseResult<CodeFieldModel[]>>('/api/autoCode/getFieldsByTableName', {
    method: 'GET',
    params: {
      tableName,
    },
  });

  return result.data || [];
}