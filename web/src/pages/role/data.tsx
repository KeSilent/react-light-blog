import { RoleModel } from '@/models/system/role-model';
import { DeleteOutlined } from '@ant-design/icons';
import { ActionType, ProColumns } from '@ant-design/pro-components';
import { Button, Popconfirm } from 'antd';
import { MutableRefObject } from 'react';
import CheckMenu from './components/CheckMenu';
import CreateRole from './components/CreateRole';

export const columns = (
  actionRef: MutableRefObject<ActionType | undefined>,
  handleDelete: (roleId: string) => void,
): ProColumns<RoleModel>[] => [
  {
    title: '角色ID',
    width: 120,
    dataIndex: 'id',
    key: 'id',
    search: false,
    hideInTable: true,
  },
  {
    title: '角色名称',
    dataIndex: 'roleName',
    key: 'roleName',
  },
  {
    title: '操作',
    width: 240,
    dataIndex: 'option',
    valueType: 'option',
    render: (_, record) => [
      <CheckMenu key="checkMenu" roleId={record.id} reload={actionRef.current?.reload} />,
      <CreateRole key="updateUser" role={record} reload={actionRef.current?.reload} />,
      <Popconfirm
        key="delete"
        title="是否确认删除角色"
        onConfirm={() => handleDelete(record.id)}
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
