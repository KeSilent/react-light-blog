export interface RoleModel {
  id: string; // 角色ID
  roleName: string; // 角色名称
  parentId: number; // 父角色ID
  createTime: string; // 创建时间
  updateTime: string | null; // 更新时间（可能为null）
  defaultRouter: string; // 默认路由
  deletedAt: string | null; // 删除时间（可能为null）
  menus: any | null; // 菜单列表（可能为null）
  users: any | null; // 用户列表（可能为null）
}

export interface RoleMenuModel {
  sysBaseMenuId: string;
  sysRoleRoleId: string;
}
