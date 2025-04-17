import { UserModel } from './user-model';

export interface DeptModel {
  id: string;
  deptName: string;
  parentId: string;
  sort: number;
  remark: string;
  status: boolean | string;
  parent: string;
  createdAt: string;
  updatedAt: string;
  deletedAt: string | null;
  users: UserModel[];
  children: DeptModel[];
}
