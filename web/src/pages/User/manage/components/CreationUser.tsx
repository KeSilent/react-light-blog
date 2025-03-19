import { UserModel } from "@/models/user-model";
import { saveUser } from "@/services/user/api";
import { PlusOutlined } from "@ant-design/icons";
import { ActionType, ModalForm, ProForm, ProFormText } from "@ant-design/pro-components";
import { FormattedMessage, useRequest } from "@umijs/max";
import { Button, Form, message } from "antd";
import { useEffect, useState } from "react";


export type CreationFormProps = {
  reload?: ActionType['reload'];
};
const CreationUser: React.FC<CreationFormProps> = (props) => {
  const { reload } = props;

  const [messageApi, contextHolder] = message.useMessage();

  const [form] = Form.useForm<UserModel>();
  const [open, setOpen] = useState(false);

  const { run, loading } = useRequest<UserModel>(saveUser, {
    manual: true,
    debounceInterval: 300,
    onSuccess: () => {
      messageApi.success('创建成功！');
      form.resetFields(); // 成功后清空表单
      reload?.(); // 刷新表格数据
    },
    onError: (error) => {
      // 处理网络错误等异常情况
      messageApi.error(error?.message || '网络异常，请稍后重试！');
    },
  })

  useEffect(() => {
    if (open) {
      form.setFieldsValue({});
    }
  }, [open, form]); // 添加 open 依赖

  // const { } = userRequest(addR)

  return (
    <>
      {contextHolder}
      <ModalForm<UserModel> title="编辑用户"
        trigger={
          <Button
            type="primary"
            key="primary"
            onClick={() => {
              setOpen(true);
            }}
          >
            <PlusOutlined /> <FormattedMessage id="pages.searchTable.new" defaultMessage="New" />
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
            form.setFieldsValue({});
          } else {
            form.resetFields();
          }
        }}
      >

        <ProForm.Group>
          <ProFormText
            label="ID"
            placeholder="id"
            width="md"
            name="id"
            hidden={true}
          />
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
              }
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
              }
            ]}
          />
        </ProForm.Group>
        <ProForm.Group>
          <ProFormText.Password
            label="密码"
            placeholder="密码"
            width="md"
            name="password"
            rules={[
              {
                required: true,
                message: '请输入新密码',
              },
              {
                min: 6,
                message: '密码长度不能小于6位',
              },
              ({ getFieldValue }) => ({
                validator(_, value) {
                  if (!value || !getFieldValue('newPassword') || value === getFieldValue('newPassword')) {
                    return Promise.resolve();
                  }
                  return Promise.reject(new Error('两次输入的密码不一致！'));
                },
              }),
            ]}
          />
          <ProFormText.Password
            label="确认新密码"
            placeholder="确认新密码"
            width="md"
            name="newPassword"
            rules={[
              {
                required: true,
                message: '请确认新密码',
              },
              ({ getFieldValue }) => ({
                validator(_, value) {
                  if (!value || !getFieldValue('password') || getFieldValue('password') === value) {
                    return Promise.resolve();
                  }
                  return Promise.reject(new Error('两次输入的密码不一致！'));
                },
              }),
            ]}
            dependencies={['newPassword']} // 添加依赖，当 newPassword 变化时重新验证
          />
        </ProForm.Group>

        <ProForm.Group>
          <ProFormText
            label="手机号"
            placeholder="手机号"
            width="md"
            name="phone"
          />
          <ProFormText
            label="邮箱"
            placeholder="邮箱"
            width="md"
            name="email"
          />
        </ProForm.Group>
      </ModalForm>
    </>
  );
};


export default CreationUser;