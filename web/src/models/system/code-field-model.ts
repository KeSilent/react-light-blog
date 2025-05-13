export interface CodeFieldModel {
  Field: string; // 字段名
  Type: string; // 字段类型
  Collation: string; // 字符集
  Null: string; // 是否允许为空
  Key: string; // 索引类型
  Default: string; // 默认值
  Extra: string; // 额外信息
  Privileges: string; // 权限
  Comment: string; // 注释
}

export interface RelationModel {
  RelateTable: string; //被关联表名
  RelateType: string; //关联类型
  RelateColumn: string; //外键字段名
  FieldName: string; //生成属性
  RelationTable: string; //新关联关系表名
}
