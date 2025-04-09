/*
 * @Author: Yang
 * @Date: 2025-04-05 11:17:58
 * @Description: 创建、编辑菜单
 */

import { MenuModel } from '@/models/system/menu-model';
import { getMenuByKey, saveMenu } from '@/services/system/menuApi';
import * as icons from '@ant-design/icons';
import {
  ActionType,
  ModalForm,
  ProForm,
  ProFormDigit,
  ProFormSelect,
  ProFormText,
  ProFormTreeSelect,
} from '@ant-design/pro-components';
import { useRequest } from '@umijs/max';
import { Button, Form, message } from 'antd';
import { useEffect, useState } from 'react';

export type CreateMenuProps = {
  menu?: MenuModel;
  reload?: ActionType['reload'];
};

export default function CreateMenu(props: CreateMenuProps) {
  const { menu, reload } = props;
  const [open, setOpen] = useState(false);
  const [messageApi, contextHolder] = message.useMessage();
  const [form] = Form.useForm<MenuModel>();

  const { run, loading } = useRequest<MenuModel>(saveMenu, {
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
      if (menu) {
        form.setFieldsValue(menu);
      } else {
        form.setFieldsValue({});
      }
    }
  }, [open, form, menu]);

  // 动态生成图标的 valueEnum
  const iconOptions = Object.keys(icons)
    .filter((key) => key.endsWith('Outlined')) // 只选择 Outlined 类型的图标
    .reduce((acc, key) => {
      const IconComponent = icons[key]; // 获取图标组件
      acc[key] = <IconComponent />; // 将图标组件作为值
      return acc;
    }, {} as Record<string, JSX.Element>);

  return (
    <>
      {contextHolder}
      <ModalForm<MenuModel>
        title={menu ? '编辑菜单' : '新建菜单'}
        trigger={
          menu ? (
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
              <icons.PlusOutlined /> 新建菜单
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
      >
        <ProForm.Group>
          <ProFormText
            label="菜单显示名称"
            placeholder="菜单显示名称"
            width="md"
            name="title"
            rules={[
              {
                required: true,
                message: '请输入菜单显示名称',
              },
            ]}
          />
          <ProFormText
            label="菜单路由名称"
            placeholder="菜单路由名称"
            width="md"
            name="name"
            rules={[
              {
                required: true,
                message: '请输入菜单路由名称',
              },
            ]}
          />
        </ProForm.Group>
        <ProForm.Group>
          <ProFormText
            label="路由显示地址"
            placeholder="例如：/system/user"
            width="md"
            name="path"
            rules={[
              {
                required: true,
                message: '请输入路由地址',
              },
            ]}
          />
          <ProFormText
            label="组件真实地址"
            placeholder="例如：./user/manage"
            width="md"
            name="component"
            rules={[
              {
                required: true,
                message: '请输入组件地址',
              },
            ]}
          />
        </ProForm.Group>

        <ProForm.Group>
          <ProFormTreeSelect
            name="parentId"
            label="上级菜单"
            width="md"
            debounceTime={300}
            fieldProps={{
              variant: 'outlined',
              bordered: undefined, 
              fieldNames: {
                label: 'title',
                value: 'uuid',
                children: 'children',
              },
            }}
            request={async () => {
              try {
                return await getMenuByKey({ keyWord: '' });
              } catch (error) {
                console.error('获取菜单数据失败:', error);
                return [];
              }
            }}
          />
          <ProFormSelect
            name="icon"
            label="图标"
            width="md"
            allowClear={false}
            showSearch
            valueEnum={iconOptions}
            fieldProps={{
              optionLabelProp: 'value',
              optionRender: (option) => (
                <div>
                  {option.label} {option.value}
                </div>
              ),
            }}
          />
        </ProForm.Group>
        <ProForm.Group>
          <ProFormSelect
            name="hidden"
            width="md"
            label="是否隐藏"
            valueEnum={{
              0: '不隐藏',
              1: '隐藏',
            }}
          />
          <ProFormDigit width="md" min={1} label="排序" name="sort" />
        </ProForm.Group>
      </ModalForm>
    </>
  );
}
