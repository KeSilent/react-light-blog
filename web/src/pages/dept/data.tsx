import { DeptModel } from '@/models/system/dept-model';
import { ActionType, ProColumns } from '@ant-design/pro-components';
import { MutableRefObject } from 'react';
import CreateDept from './components/CreateDept';
import { Button, Popconfirm } from 'antd';
import { DeleteOutlined } from '@ant-design/icons';

export const columns = (
  actionRef: MutableRefObject<ActionType | undefined>,
  handleDelete: (roleId: string) => void,
): ProColumns<DeptModel>[] => [
  {
    title: '部门名称',
    dataIndex: 'title',
    key: 'title',
  },
  {
    title: '部门状态',
    dataIndex: 'status',
    key: 'status',
  },
  {
    title: '操作',
    width: 240,
    dataIndex: 'option',
    valueType: 'option',
    render: (_, record) => [
      <CreateDept key="update" model={record} reload={actionRef.current?.reload} />,
      <Popconfirm
        key="delete"
        title="是否确认删除菜单"
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
