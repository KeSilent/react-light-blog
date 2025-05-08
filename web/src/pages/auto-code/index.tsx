import { CodeFieldModel } from '@/models/system/code-field-model';
import { getAllTableName, getFieldsByTableName } from '@/services/base/autoCodeApi';
import {
  ProFormInstance,
  ProFormSelect,
  ProTable,
  QueryFilter,
  RequestOptionsType,
} from '@ant-design/pro-components';
import { useRef } from 'react';
import { columns } from './data';

export default function AutoCode() {
  const formRef = useRef<ProFormInstance>();

  return (
    <>
      <QueryFilter
        defaultCollapsed
        split
        onFinish={async (values) => {
          console.log(values);
          await getFieldsByTableName(values.tableName);
          return true;
        }}
      >
        <ProFormSelect
          label="表"
          placeholder="表"
          width="md"
          name="tableName"
          debounceTime={300}
          request={async () => {
            const response = await getAllTableName();
            return (response || []).map((table: string) => ({
              label: table,
              value: table,
            })) as RequestOptionsType[];
          }}
          rules={[
            {
              required: true,
              message: '请选择表',
            },
          ]}
        />
      </QueryFilter>
      <ProTable<CodeFieldModel>
        columns={columns} />
    </>
  );
}
