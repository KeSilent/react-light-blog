import { CodeFieldModel } from '@/models/system/code-field-model';
import { ProColumns } from '@ant-design/pro-components';

export const columns: ProColumns<CodeFieldModel>[] = [
  {
    title: '字段名',
    dataIndex: 'Field',
    key: 'Field',
  },
  {
    title: '字段类型',
    dataIndex: 'Type',
    key: 'Type',
  },
  {
    title: '字符集',
    dataIndex: 'Collation',
    key: 'Collation',
  },
  {
    title: '是否允许为空',
    dataIndex: 'Null',
    key: 'Null',
  },
  {
    title: '索引类型',
    dataIndex: 'Key',
    key: 'Key',
  },
  {
    title: '默认值',
    dataIndex: 'Default',
    key: 'Default',
  },
  {
    title: '额外信息',
    dataIndex: 'Extra',
    key: 'Extra',
  },
  {
    title: '权限',
    dataIndex: 'Privileges',
    key: 'Privileges',
  },
  {
    title: '注释',
    dataIndex: 'Comment',
    key: 'Comment',
  },
];
