import { PageList, ResponseResult } from '@/models/system/common-model';
import { ChangePasswordReq, UserModel } from '@/models/system/user-model';
import { request } from '@umijs/max';

/**
 * 获取用户列表
 * @param params
 * @param options
 * @returns
 */
export async function getUserList(
  params: {
    current?: number;
    pageSize?: number;
  },
  options?: { [keyword: string]: any },
) {
  const result = await request<ResponseResult<PageList<UserModel[]>>>('/api/user/getUserList', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });

  return result.data;
}

export async function register(params: UserModel) {
  console.log(params);

  // const result = await request<ResponseResult>('/api/user/register', {
  //   method: 'POST',
  //   data: params,
  // });
  // return { data: result.code === 0 };
  return true;
}

export async function changePassword(params: ChangePasswordReq) {
  const result = await request<ResponseResult>('/api/user/changePassword', {
    method: 'POST',
    data: params,
  });

  return { data: result.code === 0 };
}

export async function deleteUser(params: { id: string }) {
  const result = await request('/api/user/deleteUser', {
    method: 'POST',
    data: params,
  });
  return { data: result.code === 0 };
}
export async function updateUser(params: UserModel) {
  const result = await request<ResponseResult>('/api/user/updateUser', {
    method: 'POST',
    data: params,
  });

  return { data: result.code === 0 };
}
