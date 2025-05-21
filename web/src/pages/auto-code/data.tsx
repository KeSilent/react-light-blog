import { CodeFieldModel, RelationModel } from '@/models/system/code-field-model';
import { DeleteOutlined } from '@ant-design/icons';
import { ActionType, ProColumns, RequestOptionsType } from '@ant-design/pro-components';
import { Button, Popconfirm } from 'antd';
import { MutableRefObject } from 'react';
import CreateField from './components/CreateRelation';

export const columns: ProColumns<CodeFieldModel>[] = [
  {
    title: '字段名',
    dataIndex: 'field',
    key: 'Field',
  },
  {
    title: '字段类型',
    dataIndex: 'type',
    key: 'Type',
  },
  {
    title: '字符集',
    dataIndex: 'collation',
    key: 'Collation',
  },
  {
    title: '是否允许为空',
    dataIndex: 'null',
    key: 'Null',
  },
  {
    title: '默认值',
    dataIndex: 'default',
    key: 'Default',
  },
  {
    title: '额外信息',
    dataIndex: 'extra',
    key: 'Extra',
  },
  {
    title: '注释',
    dataIndex: 'comment',
    key: 'Comment',
  },
];

export const relationColumns = (
  actionRef: MutableRefObject<ActionType | undefined>,
  handleDelete: (roleId: string) => void,
  tableData: CodeFieldModel[],
  relationData: RelationModel[],
  tableNames: RequestOptionsType[],
  setRelationData: (data: RelationModel[]) => void,
): ProColumns<RelationModel>[] => [
  {
    title: '被关联表名',
    dataIndex: 'relateTable',
    key: 'RelateTable',
  },
  {
    title: '关联类型',
    dataIndex: 'relateType',
    key: 'RelateType',
  },
  {
    title: '外键字段名',
    dataIndex: 'relateColumn',
    key: 'RelateColumn',
  },
  {
    title: '关联属性名',
    dataIndex: 'fieldName',
    key: 'FieldName',
  },
  {
    title: '新关联关系表名',
    dataIndex: 'relationTable',
    key: 'RelationTable',
  },
  {
    title: '操作',
    dataIndex: 'collation',
    key: 'Collation',
    render: (_, record) => [
      <CreateField
        key="update"
        model={record}
        reload={actionRef.current?.reload}
        codeFieldModel={tableData}
        relationModel={relationData}
        tableNames={tableNames}
        onSuccess={(data) => {
          setRelationData([...relationData, data]);
        }}
      />,
      <Popconfirm
        key="delete"
        title="是否确认删除当前数据？"
        onConfirm={() => handleDelete(record.key)}
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
