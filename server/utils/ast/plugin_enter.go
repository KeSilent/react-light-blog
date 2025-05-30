package ast

type PluginEnter struct {
	Base
	Type            Type   // 类型
	Path            string // 文件路径
	ImportPath      string // 导包路径
	RelativePath    string // 相对路径
	StructName      string // 结构体名称
	StructCamelName string // 结构体小驼峰名称
	ModuleName      string // 模块名称
	GroupName       string // 分组名称
	PackageName     string // 包名
	ServiceName     string // 服务名称
}
