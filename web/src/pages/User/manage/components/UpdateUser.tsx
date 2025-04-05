import { UserModel } from '@/models/system/user-model';
import { updateUser } from '@/services/system/userApi';
import { ActionType, ModalForm, ProForm, ProFormText } from '@ant-design/pro-components';
import { useRequest } from '@umijs/max';
import { Button, Form, message } from 'antd';
import { useEffect, useState } from 'react';

export type UpdateFormProps = {
  record?: UserModel;
  reload?: ActionType['reload'];
};
const UpdateUser: React.FC<UpdateFormProps> = (props) => {
  const { reload, record } = props;

  const [messageApi, contextHolder] = message.useMessage();

  const [form] = Form.useForm<UserModel>();
  const [open, setOpen] = useState(false);

  const { run, loading } = useRequest<UserModel>(updateUser, {
    manual: true,
    debounceInterval: 300,
    onSuccess: (res) => {
      if (res) {
        messageApi.success('修改成功！');
        form.resetFields();
        reload?.();
      }
    },
    onError: (error) => {
      messageApi.error(error?.message || '网络异常，请稍后重试！');
    },
  });

  useEffect(() => {
    if (open && record) {
      form.setFieldsValue(record);
    }
  }, [record, open, form]);

  return (
    <>
      {contextHolder}
      <ModalForm<UserModel>
        title="编辑用户"
        trigger={
          <Button color="primary" variant="link" onClick={() => setOpen(true)}>
            编辑
          </Button>
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
            form.setFieldsValue(record || {}); // 强制绑定当前数据
          } else {
            form.resetFields(); // 关闭时重置
          }
        }}
      >
        <ProForm.Group>
          <ProFormText label="ID" placeholder="id" width="md" name="id" hidden={true} />
        </ProForm.Group>
        <ProForm.Group>
          <ProFormText
            label="昵称"
            placeholder="昵称"
            width="md"
            name="nickName"
            rules={[
              {
                required: true,
                message: '请输入昵称',
              },
            ]}
          />
          <ProFormText
            label="用户名"
            placeholder="用户名"
            width="md"
            name="username"
            rules={[
              {
                required: true,
                message: '请输入用户名',
              },
            ]}
          />
        </ProForm.Group>

        <ProForm.Group>
          <ProFormText label="手机号" placeholder="手机号" width="md" name="phone" />
          <ProFormText label="邮箱" placeholder="邮箱" width="md" name="email" />
        </ProForm.Group>
      </ModalForm>
    </>
  );
};

export default UpdateUser;
