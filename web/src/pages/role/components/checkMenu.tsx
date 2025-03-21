import { MenuModel } from '@/models/system/menu-model';
import { getMenuList } from '@/services/system/menuApi';
import { getRoleMenus } from '@/services/system/roleApi';
import { changePassword } from '@/services/system/userApi';
import { ActionType, DrawerForm, ProForm } from '@ant-design/pro-components';
import { useRequest } from '@umijs/max';
import { Button, Input, message, Tree } from 'antd';
import { useEffect, useState } from 'react';

export type CheckMenuProps = {
  reload?: ActionType['reload'];
  roleId?: string;
};
export default function CheckMenu<CheckMenuProps>(props: CheckMenuProps) {
  const { reload, roleId } = props;
  const [messageApi] = message.useMessage();
  const [keyword, setKeyword] = useState('');
  const [open, setOpen] = useState(false);
  const [treeData, setTreeData] = useState<MenuModel[]>([]);
  const [selectedKeys, setSelectedKeys] = useState<string[]>([]);
  const [expandedKeys, setExpandedKeys] = useState<string[]>([]);

  const { run, loading } = useRequest<MenuModel>(changePassword, {
    manual: true,
    // 添加防抖，避免重复提交
    debounceInterval: 300,
    onSuccess: (res) => {
      if (res) {
        messageApi.success('修改成功！');
        reload?.();
      }
    },
    onError: (error) => {
      // 处理网络错误等异常情况
      messageApi.error(error?.message || '网络异常，请稍后重试！');
    },
  });

  const getSelectedKeys = () => {
    getRoleMenus(roleId).then((res) => {
      if (res) {
        const menuIdList: string[] = res.map((item) => String(item.sysBaseMenuId));
        setSelectedKeys(menuIdList);
      }
    });
  };

  useEffect(() => {
    if (open) {
      getMenuList({ keyword: '' }).then((res) => {
        if (res) {
          setTreeData(res);
          setExpandedKeys(res.map((item) => String(item.id)));
          getSelectedKeys();
        }
      });
    }
  }, [open, roleId]);

  return (
    <>
      <DrawerForm<{
        name: string;
        company: string;
      }>
        title="设置权限"
        resize={{
          maxWidth: window.innerWidth * 0.8,
          minWidth: 500,
        }}
        trigger={
          <a type="primary" onClick={() => setOpen(true)}>
            设置权限
          </a>
        }
        autoFocusFirstInput
        drawerProps={{
          destroyOnClose: true,
        }}
        submitTimeout={2000}
        onFinish={async (value) => {
          await run(value);
          return true;
        }}
        onOpenChange={(visible) => {
          setOpen(visible);
        }}
      >
        <ProForm.Group>
          <Input value={keyword} placeholder="筛选" />
          <Button type="primary">确定</Button>
        </ProForm.Group>
        <ProForm.Group>
          <Tree
            checkable
            multiple={true}
            checkedKeys={selectedKeys}
            expandedKeys={expandedKeys}
            treeData={treeData}
            fieldNames={{ title: 'title', key: 'id', children: 'children' }}
          />
        </ProForm.Group>
      </DrawerForm>
    </>
  );
}
