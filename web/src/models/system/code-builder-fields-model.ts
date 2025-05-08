export interface CodeBuilderFieldModel {
  id: string;
  fieldName: string;
  fieldDesc: string;
  fieldType: string;
  fieldJson: string;
  comment: string;
  columnName: string;
  gormTag: string;
  require: boolean;
  errorText: string;
  clearable: boolean;
  sort: boolean;
  fieldSearchType: string;
  orderDirection: string;
  dictType: string;
  structId: string;
}
