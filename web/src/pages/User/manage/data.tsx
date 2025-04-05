import { UserModel } from '@/models/system/user-model';
import { ActionType, ProColumns } from '@ant-design/pro-components';
import { Avatar, Button, Popconfirm } from 'antd';
import { MutableRefObject } from 'react';
import ChangePassword from './components/ChangePassword';
import UpdateUser from './components/UpdateUser';
import { DeleteOutlined } from '@ant-design/icons';

export const columns = (
  actionRef: MutableRefObject<ActionType | undefined>,
  handleDelete: (id: string) => void,
): ProColumns<UserModel>[] => [
  {
    title: '头像',
    width: 80,
    dataIndex: 'avatar',
    key: 'avatar',
    render: (_, record) => {
      return (
        <Avatar
          src={record.avatar}
          size="default"
          alt={record.nickName}
          style={{ cursor: 'pointer' }}
          onClick={() => window.open(record.avatar)}
        />
      );
    },
    search: false,
  },
  {
    title: 'ID',
    width: 120,
    dataIndex: 'id',
    key: 'id',
    search: false,
    hideInTable: true,
  },
  {
    title: '用户名',
    width: 120,
    dataIndex: 'username',
    key: 'username',
  },
  {
    title: '昵称',
    width: 120,
    dataIndex: 'nickName',
    key: 'nickName',
  },
  {
    title: '手机号',
    width: 120,
    dataIndex: 'phone',
    key: 'phone',
  },
  {
    title: '邮箱',
    width: 120,
    dataIndex: 'email',
    key: 'email',
  },
  {
    title: '操作',
    width: 160,
    dataIndex: 'option',
    valueType: 'option',
    render: (_, record) => [
      <UpdateUser key="updateUser" reload={actionRef.current?.reload} record={record} />,
      <ChangePassword key="changePassword" reload={actionRef.current?.reload} />,
      <Popconfirm
        key="delete"
        title="是否确认删除角色"
        onConfirm={() => handleDelete(record.uuid!)}
        okText="Yes"
        cancelText="No"
      >
        <Button color="danger" variant="link">
          <DeleteOutlined /> 删除
        </Button>
      </Popconfirm>,
    ],
  },
];
