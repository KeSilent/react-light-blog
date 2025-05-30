package request

type AutoCodeReq struct {
	PackageName        string        `json:"packageName"`
	Module             string        `json:"module"`
	TableName          string        `json:"tableName"`
	AutoCreateResource bool          `json:"autoCreateResource"` // 是否自动创建资源标识
	Relations          []RelationReq `json:"relations"`
	Template           string        `json:"template" default:"plugin"`        // 模板
	GenerateServer     bool          `json:"generateServer"`                   //  是否生成服务
	StructName         string        `json:"structName" example:"Struct名称"`    // Struct名称
	Abbreviation       string        `json:"abbreviation" example:"Struct简称"`  // Struct简称
	HumpPackageName    string        `json:"humpPackageName" example:"go文件名称"` // go文件名称
	Description        string        `json:"description"`
}

/**
 * @Author: Yang
 * @description: 关联关系
 * @return {*}
 */
type RelationReq struct {
	RelateTable   string `json:"relateTable"`
	RelateType    string `json:"relateType"`
	RelateColumn  string `json:"relateColumn"`
	RieldName     string `json:"fieldName"`
	RelationTable string `json:"relationTable"`
}

type AutoCodeField struct {
	FieldName       string `json:"fieldName"`       // Field名
	FieldDesc       string `json:"fieldDesc"`       // 中文名
	FieldType       string `json:"fieldType"`       // Field数据类型
	FieldJson       string `json:"fieldJson"`       // FieldJson
	DataTypeLong    string `json:"dataTypeLong"`    // 数据库字段长度
	Comment         string `json:"comment"`         // 数据库字段描述
	ColumnName      string `json:"columnName"`      // 数据库字段
	FieldSearchType string `json:"fieldSearchType"` // 搜索条件
	FieldSearchHide bool   `json:"fieldSearchHide"` // 是否隐藏查询条件
	DictType        string `json:"dictType"`        // 字典
	//Front           bool        `json:"front"`           // 是否前端可见
	Form            bool   `json:"form"`            // 是否前端新建/编辑
	Table           bool   `json:"table"`           // 是否前端表格列
	Desc            bool   `json:"desc"`            // 是否前端详情
	Excel           bool   `json:"excel"`           // 是否导入/导出
	Require         bool   `json:"require"`         // 是否必填
	DefaultValue    string `json:"defaultValue"`    // 是否必填
	ErrorText       string `json:"errorText"`       // 校验失败文字
	Clearable       bool   `json:"clearable"`       // 是否可清空
	Sort            bool   `json:"sort"`            // 是否增加排序
	PrimaryKey      bool   `json:"primaryKey"`      // 是否主键
	CheckDataSource bool   `json:"checkDataSource"` // 是否检查数据源
	FieldIndexType  string `json:"fieldIndexType"`  // 索引类型
}
