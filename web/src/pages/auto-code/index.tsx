import { CodeFieldModel, RelationModel } from '@/models/system/code-field-model';
import { getAllTableName, getFieldsByTableName } from '@/services/base/autoCodeApi';
import {
  ActionType,
  ProFormInstance,
  ProFormSelect,
  ProFormText,
  ProTable,
  QueryFilter,
  RequestOptionsType,
  StepsForm,
} from '@ant-design/pro-components';
import { Divider, message } from 'antd';
import { useRef, useState } from 'react';
import CreateField from './components/CreateRelation';
import { columns, relationColumns } from './data';

export default function AutoCode() {
  const formRef = useRef<ProFormInstance>();
  const ref = useRef<ActionType>();
  const relationRef = useRef<ActionType>();
  const [tableData, setTableData] = useState<CodeFieldModel[]>([]);
  const [relationData, setRelationData] = useState<RelationModel[]>([]);
  const [tableNames, setTableNames] = useState<RequestOptionsType[]>([]);
  const [messageApi, contextHolder] = message.useMessage();
  const handleDelete = async (id: string) => {
    try {
      setRelationData(relationData.filter((item) => item.key !== id));
      message.success('删除成功');
      relationRef.current?.reload();
    } catch (error) {
      message.error('删除失败');
    }
  };

  return (
    <>
      {contextHolder}
      <StepsForm<{
        name: string;
      }>
        containerStyle={{ width: '100%', maxWidth: '100%' }}
        formRef={formRef}
        onFinish={async () => {
          message.success('提交成功');
        }}
        stepsFormRender={(formDom, submitterDom) => (
          <div>
            {formDom}
            <div
              style={{
                width: '100%',
                display: 'flex',
                justifyContent: 'right',
                marginTop: 24,
                gap: 24,
              }}
            >
              {submitterDom}
            </div>
          </div>
        )}
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
            if (!tableData?.length) {
              messageApi.error('请选择表');
              return false;
            } else {
              return true;
            }
          }}
        >
          <QueryFilter
            defaultCollapsed
            split
            onFinish={async (values) => {
              const fields = await getFieldsByTableName(values.tableName);
              const dataWithKey = (fields || []).map((item) => ({
                ...item,
                key: item.field,
              }));
              setTableData(dataWithKey);
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
                const result = (response || []).map((table: string) => ({
                  label: table,
                  value: table,
                })) as RequestOptionsType[];
                setTableNames(result);
                return result;
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
            key={'codeFieldData'}
            search={false}
            columns={columns}
            dataSource={tableData}
            pagination={false}
          />
          <Divider>设置模型关系</Divider>

          <ProTable<RelationModel>
            actionRef={relationRef}
            key={'relationData'}
            toolBarRender={() => [
              <CreateField
                key="addRelation"
                reload={relationRef.current?.reload}
                codeFieldModel={tableData}
                relationModel={relationData}
                tableNames={tableNames}
                onSuccess={(data) => {
                  setRelationData([...relationData, data]);
                }}
              />,
            ]}
            search={false}
            columns={relationColumns(
              relationRef,
              handleDelete,
              tableData,
              relationData,
              tableNames,
              setRelationData,
            )}
            dataSource={relationData}
            pagination={false}
          />
        </StepsForm.StepForm>
        <StepsForm.StepForm<{
          packageName: string;
        }>
          name="checkbox"
          title="设置API"
          layout="horizontal"
          onFinish={async () => {
            console.log(formRef.current?.getFieldsValue());
            return true;
          }}
        >
          <div
            style={{
              display: 'flex',
              flexDirection: 'column',
              alignItems: 'center',
              width: '100%',
            }}
          >
            <ProFormText
              name="packageName"
              label="设置包名"
              width="md"
              rules={[{ required: true, message: '请输入包名' }]}
            />
          </div>
        </StepsForm.StepForm>
      </StepsForm>
    </>
  );
}
