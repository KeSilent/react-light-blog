import { PageParams } from "@/models/common-model";
import { UserModel } from "@/models/user-model";
import { getUserList } from "@/services/user/api";
import { PlusOutlined } from "@ant-design/icons";
import { ActionType, PageContainer, ProTable } from "@ant-design/pro-components";
import { FormattedMessage } from "@umijs/max";
import { Button } from "antd";
import { useRef, useState } from "react";
import { columns } from "./data";
import CreationUser from "./components/CreationUser";

const UserList: React.FC = () => {


  const actionRef = useRef<ActionType>();

  const handleModalOpen = (flag: boolean) => {
    console.log(flag);
  }

  const [selectedRowsState, setSelectedRows] = useState<UserModel[]>([]);

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
        request={getUserList}
        columns={columns(actionRef)}
        rowSelection={{
          onChange: (_, selectedRows) => {
            console.log(selectedRows);

            setSelectedRows(selectedRows);
          },
        }}
      />
    </PageContainer>
  )
}


export default UserList;