import { RoleModel } from '@/models/system/role-model';
import { saveRole } from '@/services/system/roleApi';
import { PlusOutlined } from '@ant-design/icons';
import { ActionType, ModalForm, ProForm, ProFormText } from '@ant-design/pro-components';
import { useRequest } from '@umijs/max';
import { Button, Form, message } from 'antd';
import { useEffect, useState } from 'react';

export type CreateRoleProps = {
  role?: RoleModel;
  reload?: ActionType['reload'];
};

export default function CreateRole(props: CreateRoleProps) {
  const { role, reload } = props;
  const [messageApi, contextHolder] = message.useMessage();
  const [open, setOpen] = useState(false);
  const [form] = Form.useForm<RoleModel>();

  const { run, loading } = useRequest<RoleModel>(saveRole, {
    manual: true,
    debounceInterval: 300,
    onSuccess: () => {
      messageApi.success('保存成功！');
      form.resetFields();
      reload?.();
    },
    onError: (error) => {
      // 处理网络错误等异常情况
      messageApi.error(error?.message || '网络异常，请稍后重试！');
    },
  });

  useEffect(() => {
    if (open) {
      if (role) {
        form.setFieldsValue(role);
      } else {
        form.setFieldsValue({});
      }
    }
  }, [open, form, role]);

  return (
    <>
      {contextHolder}
      <ModalForm<RoleModel>
        title={role ? '编辑角色' : '新建角色'}
        trigger={
          role ? (
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
              <PlusOutlined /> 新建角色
            </Button>
          )
        }
        form={form}
        layout='horizontal'
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
      >
        <ProForm.Group>
          <ProFormText label="ID" placeholder="id" width="md" name="id" hidden={true} />
        </ProForm.Group>

        <ProForm.Group>
          <ProFormText
            label="角色名称"
            placeholder="角色名称"
            width="md"
            name="roleName"
            rules={[
              {
                required: true,
                message: '请输入角色名称',
              },
            ]}
          />
        </ProForm.Group>
      </ModalForm>
    </>
  );
}
