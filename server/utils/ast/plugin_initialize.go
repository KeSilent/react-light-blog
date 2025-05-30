/*
 * @Author: Yang
 * @Date: 2025-05-23 16:49:53
 * @Description: 追加注册插件路由到主路由中
 */
package ast

import (
	"fmt"
	"go/ast"
	"io"
)

type PluginInitialize struct {
	Base
	Type         Type   // 类型
	Path         string // 文件路径
	PluginPath   string // 插件路径
	RelativePath string // 相对路径
	ImportPath   string // 导包路径
	StructName   string // 结构体名称
	PackageName  string // 包名
}

func (a *PluginInitialize) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
	if filename == "" {
		if a.RelativePath == "" {
			filename = a.PluginPath
			a.RelativePath = a.Base.RelativePath(a.PluginPath)
			return a.Base.Parse(filename, writer)
		}
		a.PluginPath = a.Base.AbsolutePath(a.RelativePath)
		filename = a.PluginPath
	}
	return a.Base.Parse(filename, writer)
}

func (a *PluginInitialize) Injection(file *ast.File) error {
	if !CheckImport(file, a.ImportPath) {
		NewImport(a.ImportPath).Injection(file)
		funcDecl := FindFunction(file, "bizPlugin")
		stmt := CreateStmt(fmt.Sprintf("PluginInit(engine, %s.Plugin)", a.PackageName))
		funcDecl.Body.List = append(funcDecl.Body.List, stmt)
	}
	return nil
}

func (a *PluginInitialize) Rollback(file *ast.File) error {
	return nil
}

func (a *PluginInitialize) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.PluginPath
	}
	return a.Base.Format(filename, writer, file)
}
