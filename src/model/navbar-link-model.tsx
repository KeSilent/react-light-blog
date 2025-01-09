export interface NavbarLinkModel {
  id: string;
  title: string;
  url: string;
  icon?: React.ReactNode;
  subItems?: NavbarLinkModel[]; // 新增子项
}
