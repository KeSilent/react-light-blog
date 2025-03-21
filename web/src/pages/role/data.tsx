import { RoleModel } from '@/models/system/role-model';
import { ActionType, ProColumns } from '@ant-design/pro-components';
import { MutableRefObject } from 'react';
import CheckMenu from './components/checkMenu';

export const columns = (
  actionRef: MutableRefObject<ActionType | undefined>,
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
    width: 120,
    dataIndex: 'roleName',
    key: 'roleName',
  },
  {
    title: '操作',
    width: 120,
    dataIndex: 'option',
    valueType: 'option',
    render: (_, record) => [
      <CheckMenu key="checkMenu" roleId={record.id} reload={actionRef.current?.reload} />,
    ],
  },
];
