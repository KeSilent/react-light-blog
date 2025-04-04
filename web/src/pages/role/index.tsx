import { PageParams } from '@/models/system/common-model';
import { RoleModel } from '@/models/system/role-model';
import { deleteRole, getRoleList } from '@/services/system/roleApi';
import { ActionType, PageContainer, ProTable } from '@ant-design/pro-components';
import { message } from 'antd';
import { useRef } from 'react';
import CreateRole from './components/CreateRole';
import { columns } from './data';

const UserList: React.FC = () => {
  const actionRef = useRef<ActionType>();

  const handleGetUserList = async (
    params: PageParams & { pageSize?: number; current?: number; keyword?: string },
  ) => {
    const result = await getRoleList(params);
    return {
      data: result?.data,
      success: true,
      total: result?.total,
    };
  };
  const handleDelete = async (roleUUId: string) => {
    try {
      if (await deleteRole(roleUUId)) {
        message.success('删除成功');
        actionRef.current?.reload();
      }
    } catch (error) {
      message.error('删除失败');
    }
  };

  return (
    <PageContainer>
      <ProTable<RoleModel, PageParams>
        rowKey="id"
        headerTitle="角色列表"
        actionRef={actionRef}
        search={{
          labelWidth: 120,
        }}
        toolBarRender={() => [
          <CreateRole key="updateUser" reload={actionRef.current?.reload} />,
        ]}
        request={handleGetUserList}
        columns={columns(actionRef, handleDelete)}
      />
    </PageContainer>
  );
};

export default UserList;
