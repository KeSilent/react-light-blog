package router

import (
	"github.com/gin-gonic/gin"
	"{{.Module}}/middleware"
)

var {{.StructName}} = new({{.Abbreviation}})

type {{.Abbreviation}} struct {}

// Init 初始化 {{.Description}} 路由信息
func (r *{{.Abbreviation}}) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
    {
        group := private.Group("{{.Abbreviation}}").Use(middleware.OperationRecord())
        group.POST("save{{.StructName}}", api{{.StructName}}.Save{{.StructName}})   // 新建{{.Description}}
        group.DELETE("delete{{.StructName}}", api{{.StructName}}.Delete{{.StructName}}) // 删除{{.Description}}
	}
	{
	    group := private.Group("{{.Abbreviation}}")
		group.GET("get{{.StructName}}List", api{{.StructName}}.Get{{.StructName}}List)  // 获取{{.Description}}列表
	}
}
