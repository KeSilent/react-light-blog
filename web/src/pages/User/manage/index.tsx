import { PageParams } from '@/models/system/common-model';
import { UserModel } from '@/models/system/user-model';
import { deleteUser, getUserList } from '@/services/system/userApi';
import { ActionType, PageContainer, ProTable } from '@ant-design/pro-components';
import { message } from 'antd';
import { useRef } from 'react';
import CreationUser from './components/CreationUser';
import { columns } from './data';

const UserList: React.FC = () => {
  const actionRef = useRef<ActionType>();

  const handleGetUserList = async (
    params: PageParams & { pageSize?: number; current?: number; keyword?: string },
  ) => {
    const result = await getUserList(params);
    return {
      data: result?.data,
      success: true,
      total: result?.total,
    };
  };

  const handleDelete = async (userId: string) => {
    try {
      if (await deleteUser({ id: userId })) {
        message.success('删除成功');
        actionRef.current?.reload();
      }
    } catch (error) {
      message.error('删除失败');
    }
  };

  return (
    <PageContainer>
      <ProTable<UserModel, PageParams>
        rowKey="id"
        headerTitle="用户列表"
        actionRef={actionRef}
        search={{
          labelWidth: 120,
        }}
        toolBarRender={() => [<CreationUser key="updateUser" reload={actionRef.current?.reload} />]}
        request={handleGetUserList}
        columns={columns(actionRef, handleDelete)}
      />
    </PageContainer>
  );
};

export default UserList;
