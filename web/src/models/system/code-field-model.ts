export interface CodeFieldModel {
  field: string; // 字段名
  type: string; // 字段类型
  collation: string; // 字符集
  null: string; // 是否允许为空
  key: string; // 索引类型
  default: string; // 默认值
  extra: string; // 额外信息
  privileges: string; // 权限
  comment: string; // 注释
}

export interface RelationModel {
  key: string;
  relateTable: string; //被关联表名
  relateType: string; //关联类型
  relateColumn: string; //外键字段名
  fieldName: string; //生成属性
  relationTable: string; //新关联关系表名
}
