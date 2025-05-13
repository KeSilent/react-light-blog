import { CodeFieldModel, RelationModel } from '@/models/system/code-field-model';
import { getAllTableName, getFieldsByTableName } from '@/services/base/autoCodeApi';
import {
  ActionType,
  ProFormInstance,
  ProFormSelect,
  ProTable,
  QueryFilter,
  RequestOptionsType,
  StepsForm,
} from '@ant-design/pro-components';
import { Divider, message } from 'antd';
import { useRef, useState } from 'react';
import CreateField from './components/CreateField';
import { columns, relationColumns } from './data';

export default function AutoCode() {
  const formRef = useRef<ProFormInstance>();
  const ref = useRef<ActionType>();
  const relationRef = useRef<ActionType>();
  const [tableData, setTableData] = useState<CodeFieldModel[]>([]);
  const [relationData, setRelationData] = useState<RelationModel[]>([]);
  return (
    <>
      <StepsForm<{
        name: string;
      }>
        containerStyle={{ width: '100%', maxWidth: '100%' }}
        formRef={formRef}
        onFinish={async () => {
          message.success('提交成功');
        }}
        submitter={{
          render: (props, dom) => {
            return (
              <div
                style={{ width: '100%', display: 'flex', justifyContent: 'center', marginTop: 24 }}
              >
                {dom}
              </div>
            );
          },
        }}
        formProps={{
          validateMessages: {
            required: '此项为必填项',
          },
        }}
      >
        <StepsForm.StepForm<{
          name: string;
        }>
          name="one"
          title="设置模型及关系"
          layout="horizontal"
          onFinish={async () => {
            console.log(formRef.current?.getFieldsValue());
            return true;
          }}
        >
          <QueryFilter
            defaultCollapsed
            split
            onFinish={async (values) => {
              console.log(values);
              const fields = await getFieldsByTableName(values.tableName);
              setTableData(fields);
              return true;
            }}
          >
            <ProFormSelect
              label="表"
              placeholder="表"
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
            actionRef={ref}
            search={false}
            columns={columns}
            dataSource={tableData}
            pagination={false}
          />
          <Divider>设置模型关系</Divider>

          <ProTable<RelationModel>
            actionRef={relationRef}
            toolBarRender={() => [
              <CreateField key="addRelation" reload={relationRef.current?.reload} />,
            ]}
            search={false}
            columns={relationColumns}
            dataSource={relationData}
            pagination={false}
          />
        </StepsForm.StepForm>
        <StepsForm.StepForm<{
          checkbox: string;
        }>
          name="checkbox"
          title="设置参数"
          stepProps={{
            description: '这里填入运维参数',
          }}
          onFinish={async () => {
            console.log(formRef.current?.getFieldsValue());
            return true;
          }}
        ></StepsForm.StepForm>
      </StepsForm>
    </>
  );
}
