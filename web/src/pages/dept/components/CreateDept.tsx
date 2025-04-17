import { StatusConstant } from '@/constant/Status-Constant';
import { DeptModel } from '@/models/system/dept-model';
import { getListByTreeSelect, saveDept } from '@/services/system/deptApi';
import { PlusOutlined } from '@ant-design/icons';
import {
  ActionType,
  ModalForm,
  ProForm,
  ProFormDigit,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,
  ProFormTreeSelect,
} from '@ant-design/pro-components';
import { useRequest } from '@umijs/max';
import { Button, Form, message } from 'antd';
import { useEffect, useState } from 'react';

export type CreateDeptProps = {
  model?: DeptModel;
  reload?: ActionType['reload'];
};
export default function CreateDept(props: CreateDeptProps) {
  const { model, reload } = props;
  const [open, setOpen] = useState(false);
  const [messageApi, contextHolder] = message.useMessage();
  const [form] = Form.useForm<DeptModel>();

  const { run, loading } = useRequest<DeptModel>(saveDept, {
    manual: true,
    debounceInterval: 300,
    onSuccess: (res) => {
      if (!res) {
        messageApi.success('保存成功！');
        form.resetFields();
        reload?.();
      } else {
        messageApi.error('保存失败！');
      }
    },
    onError: (error) => {
      // 处理网络错误等异常情况
      messageApi.error(error?.message || '网络异常，请稍后重试！');
    },
  });

  useEffect(() => {
    if (open) {
      if (model) {
        form.setFieldsValue({
          ...model,
          status: model.status ? 'true' : 'false',
        });
      } else {
        form.setFieldsValue({
          parentId: '0',
          deptName: '',
          sort: 1,
          remark: '',
          status: 'true',
          parent: '',
          createdAt: '',
          updatedAt: '',
          deletedAt: null,
          users: [],
          children: [],
        });
      }
    }
  }, [open, form, model]);
  return (
    <>
      {contextHolder}
      <ModalForm<DeptModel>
        title={model ? '编辑部门' : '新建部门'}
        trigger={
          model ? (
            <Button color="primary" variant="link" onClick={() => setOpen(true)}>
              编辑
            </Button>
          ) : (
            <Button
              type="primary"
              key="primary"
              onClick={() => {
                setOpen(true);
              }}
            >
              <PlusOutlined /> 新建部门
            </Button>
          )
        }
        form={form}
        autoFocusFirstInput
        modalProps={{ okButtonProps: { loading }, destroyOnClose: true }}
        onFinish={async (value) => {
          await run(value);
          return true;
        }}
        onOpenChange={(visible) => {
          setOpen(visible);
          if (visible) {
            form.setFieldsValue({
              parentId: '0',
              deptName: '',
              sort: 1,
              remark: '',
              status: 'true',
              parent: '',
              createdAt: '',
              updatedAt: '',
              deletedAt: null,
              users: [],
              children: [],
            });
          } else {
            form.resetFields();
          }
        }}
      >
        <ProForm.Group>
          <ProFormText label="id" width="md" name="id" hidden={true} />
        </ProForm.Group>
        <ProForm.Group>
          <ProFormText
            name="deptName"
            label="部门名称"
            width="md"
            rules={[{ required: true, message: '请输入部门名称' }]}
          />
          <ProFormTreeSelect
            name="parentId"
            label="上级部门"
            width="md"
            debounceTime={300}
            fieldProps={{
              fieldNames: {
                label: 'deptName',
                value: 'id',
                children: 'children',
              },
            }}
            request={async () => {
              try {
                return await getListByTreeSelect({ keyWord: '' });
              } catch (error) {
                console.error('获取数据失败:', error);
                return [];
              }
            }}
          />
        </ProForm.Group>
        <ProForm.Group>
          <ProFormSelect
            name="status"
            label="部门状态"
            width="md"
            valueEnum={StatusConstant}
            transform={(value) => ({
              status: value === 'true',
            })}
          />
          <ProFormDigit name="sort" label="排序" width="md" min={1} />
        </ProForm.Group>
        <ProFormTextArea width="xl" label="备注" name="remark" />
      </ModalForm>
    </>
  );
}
