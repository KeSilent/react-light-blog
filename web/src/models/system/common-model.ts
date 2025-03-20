//分页参数
export interface PageParams {
  page?: number;
  pageSize?: number;
}

export interface PageList<T> {
  data?: T;
  /** 列表的内容总数 */
  total?: number;
  success?: boolean;
};

export interface ResponseResult<T = any> {
  code: number;
  msg: string;
  data?: T;
}