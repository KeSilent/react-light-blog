import { UserModel } from "@/models/user-model";
import { ActionType, ProColumns } from "@ant-design/pro-components";
import { MutableRefObject } from "react";

export const columns = (actionRef: MutableRefObject<ActionType | undefined>): ProColumns<UserModel>[] => [
  {
    title: '角色ID',
    width: 120,
    dataIndex: 'id',
    key: 'id',
    search: false,
    hideInTable: true
  },
  {
    title: '角色名称',
    width: 120,
    dataIndex: 'username',
    key: 'username',
  },
  {
    title: '操作',
    width: 120,
    dataIndex: 'option',
    valueType: 'option',
    render: (_, record) => [
      
    ],
  },
]