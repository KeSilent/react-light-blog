import { PageParams } from "@/models/system/common-model";
import { ActionType, PageContainer, ProTable } from "@ant-design/pro-components";
import { useRef } from "react";
import { columns } from "./data";
import { getRoleList } from "@/services/system/roleApi";
import { RoleModel } from "@/models/system/role-model";

const UserList: React.FC = () => {


  const actionRef = useRef<ActionType>();

  const handleGetUserList = async (params: PageParams & { pageSize?: number; current?: number; keyword?: string }) => {
    const result = await getRoleList(params);
    return {
      data: result?.data,
      success: true,
      total: result?.total,
    };
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
          
        ]}
        request={handleGetUserList}
        columns={columns(actionRef)}
      />
    </PageContainer>
  )
}


export default UserList;