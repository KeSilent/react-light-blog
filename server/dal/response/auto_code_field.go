package response

type AutoCodeField struct {
	Field      string `gorm:"column:Field"`      // 字段名
	Type       string `gorm:"column:Type"`       // 字段类型
	Collation  string `gorm:"column:Collation"`  // 字符集
	Null       string `gorm:"column:Null"`       // 是否允许为空
	Key        string `gorm:"column:Key"`        // 索引类型
	Default    string `gorm:"column:Default"`    // 默认值
	Extra      string `gorm:"column:Extra"`      // 额外信息
	Privileges string `gorm:"column:Privileges"` // 权限
	Comment    string `gorm:"column:Comment"`    // 注释
}
