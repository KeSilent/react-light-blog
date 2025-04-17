import { StatusOptions } from '@/constant/Status-Constant';
import { DeptModel } from '@/models/system/dept-model';
import { CheckCircleOutlined, CloseCircleOutlined, DeleteOutlined } from '@ant-design/icons';
import { ActionType, ProColumns } from '@ant-design/pro-components';
import { Button, Popconfirm, Select, Tag } from 'antd';
import { MutableRefObject } from 'react';
import CreateDept from './components/CreateDept';

export const columns = (
  actionRef: MutableRefObject<ActionType | undefined>,
  handleDelete: (roleId: string) => void,
): ProColumns<DeptModel>[] => [
  {
    title: '部门名称',
    dataIndex: 'deptName',
    key: 'deptName',
  },
  {
    title: '部门状态',
    dataIndex: 'status',
    key: 'status',
    renderFormItem: (item, { ...rest }) => {
      return <Select {...rest} allowClear={true} options={StatusOptions} />;
    },
    render: (_, record) => {
      if (record.status) {
        return (
          <Tag icon={<CheckCircleOutlined />} color="success">
            正常
          </Tag>
        );
      } else {
        return (
          <Tag icon={<CloseCircleOutlined />} color="error">
            禁用
          </Tag>
        );
      }
    },
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
        title="是否确认删除当前数据,及其包含下级数据？"
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
