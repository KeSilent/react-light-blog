import { PageParams } from '@/models/system/common-model';
import { DeptModel } from '@/models/system/dept-model';
import { deleteDept, getDeptList } from '@/services/system/deptApi';
import { ActionType, PageContainer, ProTable } from '@ant-design/pro-components';
import { useRef } from 'react';
import { columns } from './data';
import { message } from 'antd';
import CreateDept from './components/CreateDept';

export default function Dept() {
  const actionRef = useRef<ActionType>();
  const handleGetUserList = async (
    params: PageParams & { pageSize?: number; current?: number; name?: string },
  ) => {
    const result = await getDeptList(params);
    return {
      data: result?.data,
      success: true,
      total: result?.total,
    };
  };

  const handleDelete = async (id: string) => {
    try {
      if (!(await deleteDept(id))) {
        message.success('删除成功');
        actionRef.current?.reload();
      }
    } catch (error) {
      message.error('删除失败');
    }
  };

  return (
    <PageContainer>
      <ProTable<DeptModel, PageParams>
        rowKey="id"
        headerTitle="部门列表"
        actionRef={actionRef}
        search={{
          labelWidth: 120,
        }}
        toolBarRender={() => [<CreateDept key="addMenu" reload={actionRef.current?.reload} />]}
        request={handleGetUserList}
        columns={columns(actionRef, handleDelete)}
      />
    </PageContainer>
  );
}
