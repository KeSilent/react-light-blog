import { ResponseResult } from '@/models/system/common-model';
import { ChangePasswordReq } from '@/models/system/user-model';
import { changePassword } from '@/services/system/userApi';
import { ActionType, ModalForm, ProForm, ProFormText } from '@ant-design/pro-components';
import { useRequest } from '@umijs/max';
import { Button, Form, message } from 'antd';

export type UpdateFormProps = {
  reload?: ActionType['reload'];
};
const ChangePassword: React.FC<UpdateFormProps> = (props) => {
  const { reload } = props;
  const [messageApi, contextHolder] = message.useMessage();
  const [form] = Form.useForm<ChangePasswordReq>();

  const { run, loading } = useRequest<ResponseResult>(changePassword, {
    manual: true,
    // 添加防抖，避免重复提交
    debounceInterval: 300,
    onSuccess: (res) => {
      if (res) {
        messageApi.success('修改成功！');
        form.resetFields();
        reload?.();
      }
    },
    onError: (error) => {
      // 处理网络错误等异常情况
      messageApi.error(error?.message || '网络异常，请稍后重试！');
    },
  });

  return (
    <>
      {contextHolder}
      <ModalForm<ChangePasswordReq>
        title="修改密码"
        trigger={
          <Button color="primary" variant="link">
            修改密码
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
          if (!visible) {
            form.resetFields();
          }
        }}
      >
        <ProForm.Group>
          <ProFormText.Password
            label="原密码"
            placeholder={'原密码'}
            rules={[
              {
                required: true,
                message: '请填写原密码',
              },
            ]}
            width="md"
            name="password"
          />
        </ProForm.Group>

        <ProForm.Group>
          <ProFormText.Password
            label="新密码"
            placeholder="新密码"
            width="md"
            name="newPassword"
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
                  if (
                    !value ||
                    !getFieldValue('newPassword2') ||
                    value === getFieldValue('newPassword2')
                  ) {
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
            name="newPassword2"
            rules={[
              {
                required: true,
                message: '请确认新密码',
              },
              ({ getFieldValue }) => ({
                validator(_, value) {
                  if (
                    !value ||
                    !getFieldValue('newPassword') ||
                    getFieldValue('newPassword') === value
                  ) {
                    return Promise.resolve();
                  }
                  return Promise.reject(new Error('两次输入的密码不一致！'));
                },
              }),
            ]}
            dependencies={['newPassword']} // 添加依赖，当 newPassword 变化时重新验证
          />
        </ProForm.Group>
      </ModalForm>
    </>
  );
};

export default ChangePassword;
