export interface NavbarLinkModel {
  id: string;
  title: string;
  url: string;
  icon?: React.ReactNode;
  type: "link" | "group"; // 新增类型
  subItems?: NavbarLinkModel[]; // 新增子项
}
