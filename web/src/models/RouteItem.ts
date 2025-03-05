// src/models/RouteItem.ts
export interface RouteItem {
  id: number;
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
  children: RouteItem[];
}