import { MenuModel } from '@/models/system/menu-model';
import { RoleMenuModel } from '@/models/system/role-model';
import { getMenuList } from '@/services/system/menuApi';
import { addRoleMenu, getRoleMenus } from '@/services/system/roleApi';
import { ActionType, DrawerForm, ProForm } from '@ant-design/pro-components';
import { useRequest } from '@umijs/max';
import { Button, Input, message, Tree } from 'antd';
import { useEffect, useState } from 'react';

export type CheckMenuProps = {
  reload?: ActionType['reload'];
  roleId?: string;
};
export default function CheckMenu(props: CheckMenuProps) {
  const { reload, roleId } = props;
  const [messageApi, contextHolder] = message.useMessage();
  const [keyword, setKeyword] = useState('');
  const [open, setOpen] = useState(false);
  const [treeData, setTreeData] = useState<MenuModel[]>([]);
  const [selectedKeys, setSelectedKeys] = useState<string[]>([]);
  const [expandedKeys, setExpandedKeys] = useState<string[]>([]);

  const { run } = useRequest<MenuModel>(addRoleMenu, {
    manual: true,
    debounceInterval: 300,
    onSuccess: (res) => {
      if (res) {
        messageApi.success('设置成功！');
        reload?.();
      }
    },
    onError: (error) => {
      messageApi.error(error?.message || '网络异常，请稍后重试！');
    },
  });

  const getSelectedKeys = () => {
    console.log('roleId', roleId);

    if (!roleId) return;
    getRoleMenus(roleId).then((res) => {
      if (res) {
        const menuIdList: string[] = res.map((item) => String(item.sysBaseMenuUuid));
        setSelectedKeys(menuIdList);
      }
    });
  };

  const getMenus = () => {
    getMenuList({ keyWord: keyword }).then((res) => {
      if (res) {
        setTreeData(res);
        setExpandedKeys(res.map((item) => String(item.uuid)));
        getSelectedKeys();
      }
    });
  };

  useEffect(() => {
    if (open && roleId) {
      getMenus();
    }
  }, [open, roleId]);

  //输入搜索关键词
  function handleKeywordChange(e: React.ChangeEvent<HTMLInputElement>): void {
    setKeyword(e.target.value);
  }
  // 选择菜单
  const onCheckSelect = (selectedKeys: any) => {
    setSelectedKeys(selectedKeys);
  };

  return (
    <>
      {contextHolder}
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
          <Button color="primary" variant="link" onClick={() => setOpen(true)}>
            设置权限
          </Button>
        }
        autoFocusFirstInput
        drawerProps={{
          destroyOnClose: true,
        }}
        submitTimeout={2000}
        onFinish={async () => {
          let roleMenus: RoleMenuModel[] = [];
          if (!selectedKeys) {
            roleMenus.push({
              sysBaseMenuUuid: '0',
              sysRoleUuid: roleId || '',
            });
          } else {
            roleMenus = selectedKeys.map((item) => {
              return {
                sysBaseMenuUuid: item,
                sysRoleUuid: roleId || '',
              };
            });
          }
          await run({ rolemenus: roleMenus, roleUUID: roleId });
          return true;
        }}
        onOpenChange={(visible) => {
          setOpen(visible);
        }}
      >
        <ProForm.Group>
          <Input value={keyword} onChange={handleKeywordChange} placeholder="筛选" />
          <Button type="primary" onClick={getMenus}>
            确定
          </Button>
        </ProForm.Group>
        <ProForm.Group>
          <Tree
            checkable
            selectable={false}
            multiple={true}
            checkedKeys={selectedKeys}
            expandedKeys={expandedKeys}
            treeData={treeData}
            onCheck={onCheckSelect}
            onExpand={(keys) => setExpandedKeys(keys.map((key) => String(key)))}
            fieldNames={{ title: 'title', key: 'uuid', children: 'children' }}
          />
        </ProForm.Group>
      </DrawerForm>
    </>
  );
}
