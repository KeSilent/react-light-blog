import { RequestConfig } from '@umijs/max';
import { message } from 'antd';
import { history } from 'umi';

// 错误处理方案：错误类型
enum ErrorShowType {
  SILENT = 0,
  WARN_MESSAGE = 1,
  ERROR_MESSAGE = 2,
  NOTIFICATION = 3,
  REDIRECT = 9,
}

// 与后端约定的响应数据格式
interface ResponseStructure {
  success: boolean;
  data: any;
  errorCode?: number;
  errorMessage?: string;
  showType?: ErrorShowType;
}

export const requestConfig: RequestConfig = {
  // 请求拦截器
  requestInterceptors: [
    (config: any) => {
      // 拦截请求配置，进行个性化处理。
      const token = localStorage.getItem('token');
      if (token) {
        config.headers['Authorization'] = `Bearer ${token}`;
      }
      return config;
    },
  ],

  // 响应拦截器
  responseInterceptors: [
    (response) => {
      // 拦截响应数据，进行个性化处理
      const { data } = response as unknown as ResponseStructure;
      if (data.code) {
        message.error(data.errorMessage);
      }
      return response;
    },
  ],

  // 统一的错误处理
  errorConfig: {
    errorHandler: (error: any) => {
      const { response } = error;

      if (response?.status === 401) {
        message.error('登录已过期，请重新登录');
        history.push('/user/login');
      }

      if (response?.status === 500) {
        message.error('服务器错误，请稍后再试');
        history.push('/user/login');
      }
    },
  },
};
