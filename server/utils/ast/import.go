package ast

type Import struct {
	Base
	ImportPath string // 导包路径
}

func NewImport(importPath string) *Import {
	return &Import{ImportPath: importPath}
}
