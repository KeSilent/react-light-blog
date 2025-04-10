/*
 * @Author: Yang
 * @Date: 2025-04-05 08:48:10
 * @Description: 菜单列表
 */
import { PageParams } from '@/models/system/common-model';
import { MenuModel } from '@/models/system/menu-model';
import { deleteMenu, getMenuList } from '@/services/system/menuApi';
import { ActionType, PageContainer, ProTable } from '@ant-design/pro-components';
import { message } from 'antd';
import { useRef } from 'react';
import CreateMenu from './components/CreateMenu';
import { columns } from './data';

export default function Menu() {
  const actionRef = useRef<ActionType>();

  const handleGetUserList = async (
    params: PageParams & { pageSize?: number; current?: number; name?: string },
  ) => {
    const result = await getMenuList(params);
    return {
      data: result?.data,
      success: true,
      total: result?.total,
    };
  };

  const handleDelete = async (roleUUId: string) => {
    try {
      if (!(await deleteMenu(roleUUId))) {
        message.success('删除成功');
        actionRef.current?.reload();
      }
    } catch (error) {
      message.error('删除失败');
    }
  };
  return (
    <PageContainer>
      <ProTable<MenuModel, PageParams>
        rowKey="id"
        headerTitle="菜单列表"
        actionRef={actionRef}
        search={{
          labelWidth: 120,
        }}
        toolBarRender={() => [<CreateMenu key="addMenu" reload={actionRef.current?.reload} />]}
        request={handleGetUserList}
        columns={columns(actionRef, handleDelete)}
      />
    </PageContainer>
  );
}
