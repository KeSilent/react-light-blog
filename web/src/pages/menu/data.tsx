import { ProColumns } from "@ant-design/pro-components";

import { MenuModel } from '@/models/system/menu-model';

export const columns = (
  actionRef: MutableRefObject<ActionType | undefined>,
  handleDelete: (roleId: string) => void,
): ProColumns<MenuModel>[] => [
  {
    title: '菜单显示名称',
    dataIndex: 'title',
    key: 'title',
    width: '20%',
  },
  {
    title: '菜单路由名称',
    dataIndex: 'name',
    key: 'name',
    width: '20%',
  },
  {
    title: '图标',
    dataIndex: 'icon',
    key: 'icon',
    width: '20%',
  },
  {
    title: '路由地址',
    dataIndex: 'path',
    key: 'path',
    width: '20%',
  },
  {
    title: '组件地址',
    dataIndex: 'component',
    key: 'component',
    width: '20%',
  },
];
