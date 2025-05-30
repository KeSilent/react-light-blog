/*
 * @Author: Yang
 * @Date: 2025-04-22 09:39:37
 * @Description: 增加字段
 */
import { CodeFieldModel, RelationModel } from '@/models/system/code-field-model';
import { getFieldsByTableName } from '@/services/base/autoCodeApi';
import * as icons from '@ant-design/icons';
import {
  ActionType,
  ModalForm,
  ProFormSelect,
  ProFormText,
  RequestOptionsType,
} from '@ant-design/pro-components';
import { Button, Form, message } from 'antd';
import { useEffect, useState } from 'react';

export type CreateRelationProps = {
  model?: RelationModel;
  codeFieldModel?: CodeFieldModel[];
  relationModel?: RelationModel[];
  tableNames?: RequestOptionsType[];
  reload?: ActionType['reload'];
  onSuccess?: (data: RelationModel) => void;
};
export default function CreateRelation(props: CreateRelationProps) {
  const { codeFieldModel, relationModel, model, tableNames, onSuccess, reload } = props;
  const [open, setOpen] = useState(false);
  const [messageApi, contextHolder] = message.useMessage();
  const [form] = Form.useForm<RelationModel>();
  const [relateType, setRelateType] = useState('');
  const [tableFields, setTableFields] = useState<RequestOptionsType[]>([]);

  useEffect(() => {
    if (open) {
      if (model) {
        getFieldsByTableName(model.relateTable).then((res) => {
          const result = (res || []).map((field: CodeFieldModel) => ({
            label: field.field,
            value: field.field,
          })) as RequestOptionsType[];
          setTableFields(result);
          form.setFieldsValue(model);
        });
      } else {
        form.setFieldsValue({
          key: '',
          relateTable: '', //被关联表名
          relateType: '', //关联类型
          relateColumn: '', //外键字段名
          fieldName: '', //生成属性
          relationTable: '', //新关联关系表名
        });
      }
    }
  }, [open, form, model]);

  return (
    <>
      {contextHolder}
      {model ? (
        <Button color="primary" variant="link" onClick={() => setOpen(true)}>
          编辑
        </Button>
      ) : (
        <Button
          type="primary"
          key="primary"
          onClick={() => {
            if (!codeFieldModel?.length) {
              messageApi.error('请选择需要设置关系的表');
            } else {
              setOpen(true);
            }
          }}
        >
          <icons.PlusOutlined /> 新建模型关系
        </Button>
      )}
      <ModalForm<RelationModel>
        title={model ? '编辑模型关系' : '新建模型关系'}
        open={open}
        form={form}
        layout="horizontal"
        labelCol={{ span: 4 }}
        onFinish={async (value) => {
          console.log(value);
          
          if (relationModel?.find((item) => item.fieldName === value.fieldName)) {
            messageApi.error('关联属性名已存在！');
          } else {
            messageApi.success('保存成功！');
            form.resetFields();
            reload?.();
            value.key = value.fieldName;
            onSuccess?.(value);
            setOpen(false);
          }
        }}
        onOpenChange={(visible) => {
          setOpen(visible);
          if (visible) {
            form.setFieldsValue({});
          } else {
            form.resetFields();
          }
        }}
      >
        <ProFormSelect
          name="relateTable"
          label="关联表名"
          options={tableNames}
          rules={[{ required: true, message: '关联表名' }]}
          onChange={async (value: any) => {
            const fields = await getFieldsByTableName(value);
            const result = (fields || []).map((field: CodeFieldModel) => ({
              label: field.field,
              value: field.field,
            })) as RequestOptionsType[];
            setTableFields(result);
            return true;
          }}
        />
        <ProFormSelect
          name="relateType"
          label="关联类型"
          rules={[{ required: true, message: '请选择关联类型' }]}
          onChange={(value: any) => {
            setRelateType(value);
          }}
          options={['HasOne', 'HasMany', 'BelongsTo', 'Many2Many']}
        />
        <ProFormSelect
          name="relateColumn"
          label="外键字段名"
          options={tableFields}
          rules={[{ required: true, message: '请输入外键字段名' }]}
        />
        <ProFormText
          name="fieldName"
          label="关联属性名"
          rules={[{ required: true, message: '请输入关联属性名' }]}
        />
        <ProFormText
          name="relationTable"
          label="关联关系表名"
          hidden={relateType !== 'BelongsTo'}
          rules={[
            {
              required: relateType === 'BelongsTo',
              message: '请输入关联关系表名',
            },
          ]}
        />
      </ModalForm>
    </>
  );
}
