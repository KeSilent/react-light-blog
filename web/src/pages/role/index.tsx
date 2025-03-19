import { PageParams } from "@/models/common-model";
import { UserModel } from "@/models/user-model";
import { getUserList } from "@/services/user/api";
import { ActionType, PageContainer, ProTable } from "@ant-design/pro-components";
import { useRef } from "react";
import { columns } from "./data";

const UserList: React.FC = () => {


  const actionRef = useRef<ActionType>();

  const handleGetUserList = async (params: PageParams & { pageSize?: number; current?: number; keyword?: string }) => {
    const result = await getUserList(params);
    return {
      data: result?.data,
      success: true,
      total: result?.total,
    };
  };

  return (
    <PageContainer>
      <ProTable<UserModel, PageParams>
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