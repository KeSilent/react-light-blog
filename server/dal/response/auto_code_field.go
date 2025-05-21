package response

type AutoCodeField struct {
	Field      string `json:"field"`      // 字段名
	Type       string `json:"type"`       // 字段类型
	Collation  string `json:"collation"`  // 字符集
	Null       string `json:"null"`       // 是否允许为空
	Key        string `json:"key"`        // 索引类型
	Default    string `json:"default"`    // 默认值
	Extra      string `json:"extra"`      // 额外信息
	Privileges string `json:"privileges"` // 权限
	Comment    string `json:"comment"`    // 注释
}
