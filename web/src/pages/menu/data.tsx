import { ActionType, ProColumns } from '@ant-design/pro-components';

import { MenuModel } from '@/models/system/menu-model';
import { DeleteOutlined } from '@ant-design/icons';
import { Button, Popconfirm } from 'antd';
import { MutableRefObject } from 'react';
import CreateMenu from './components/CreateMenu';

export const columns = (
  actionRef: MutableRefObject<ActionType | undefined>,
  handleDelete: (roleId: string) => void,
): ProColumns<MenuModel>[] => [
  {
    title: '菜单显示名称',
    dataIndex: 'title',
    key: 'title',
  },
  {
    title: '菜单路由名称',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '图标',
    dataIndex: 'icon',
    key: 'icon',
    search: false,
  },
  {
    title: '路由地址',
    dataIndex: 'path',
    key: 'path',
  },
  {
    title: '组件地址',
    dataIndex: 'component',
    key: 'component',
  },
  {
    title: '操作',
    width: 240,
    dataIndex: 'option',
    valueType: 'option',
    render: (_, record) => [
      <CreateMenu key="updateUser" menu={record} reload={actionRef.current?.reload} />,
      <Popconfirm
        key="delete"
        title="是否确认删除菜单"
        onConfirm={() => handleDelete(record.uuid)}
        okType="danger"
        okText="删除"
        cancelText="取消"
      >
        <Button color="danger" variant="link">
          <DeleteOutlined /> 删除
        </Button>
      </Popconfirm>,
    ],
  },
];
