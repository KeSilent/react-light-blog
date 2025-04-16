import { DeptModel } from '@/models/system/dept-model';
import { saveDept } from '@/services/system/deptApi';
import { PlusOutlined } from '@ant-design/icons';
import { ActionType, ModalForm } from '@ant-design/pro-components';
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
        form.setFieldsValue(model);
      } else {
        form.setFieldsValue({});
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
        modalProps={{ okButtonProps: { loading } }}
        onFinish={async (value) => {
          await run(value);
          return true;
        }}
        onOpenChange={(visible) => {
          setOpen(visible);
          if (visible) {
            form.setFieldsValue({});
          } else {
            form.resetFields();
          }
        }}
      ></ModalForm>
    </>
  );
}
