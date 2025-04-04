// src/models/MenuItem.ts
export interface MenuModel {
  id: number;
  uuid: string;
  menuLevel: number;
  parentId: number;
  path: string;
  name: string;
  hidden: boolean;
  component: string;
  sort: number;
  title: string;
  icon: string;
  createTime: string;
  updateTime: string | null;
  deletedAt: string | null;
  authorities: any[] | null;
  children: MenuModel[];
}