import { PageParams } from "@/models/system/common-model";
import { UserModel } from "@/models/system/user-model";
import { getUserList } from "@/services/system/userApi";
import { ActionType, PageContainer, ProTable } from "@ant-design/pro-components";
import { useRef } from "react";
import { columns } from "./data";
import CreationUser from "./components/CreationUser";

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
        headerTitle="用户列表"
        actionRef={actionRef}
        search={{
          labelWidth: 120,
        }}
        toolBarRender={() => [
          <CreationUser
            key="updateUser"
            reload={actionRef.current?.reload}
          />,
        ]}
        request={handleGetUserList}
        columns={columns(actionRef)}
      />
    </PageContainer>
  )
}


export default UserList;