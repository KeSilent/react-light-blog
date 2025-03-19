import { UserModel } from "@/models/user-model";
import { ActionType, ProColumns } from "@ant-design/pro-components";
import { Avatar } from 'antd';
import { MutableRefObject } from "react";
import UpdateUser from "./components/UpdateUser";
import ChangePassword from "./components/ChangePassword";

export const columns = (actionRef: MutableRefObject<ActionType | undefined>): ProColumns<UserModel>[] => [
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
          style={{ cursor: 'pointer' }
          }
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
    hideInTable: true
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
    width: 120,
    dataIndex: 'option',
    valueType: 'option',
    render: (_, record) => [
      <UpdateUser
        key="updateUser"
        reload={actionRef.current?.reload}
        record={record}
      />,
      <ChangePassword key="changePassword" reload={actionRef.current?.reload} />
      ,
    ],
  },
]