package api

import (
	"fmt"
    
	"github.com/gin-gonic/gin"
	"{{.Module}}/dal/common/response"
	"{{.Module}}/plugin/{{.PackageName}}/model"
	req "{{.Module}}/plugin/{{.PackageName}}/request"
	"{{.Module}}/global"
	"{{.Module}}/utils"
	"go.uber.org/zap"
)

var {{.StructName}}API = new({{.Abbreviation}}Api)

type {{.Abbreviation}}Api struct {}

// Get{{.StructName}}List 分页获取{{.Description}}列表
// @Tags {{.StructName}}
// @Summary 分页获取{{.Description}}列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /{{.Abbreviation}}/get{{.StructName}}List [get]
func ({{.Abbreviation}} *{{.Abbreviation}}Api) Get{{.StructName}}ListByPage(c *gin.Context) {
	var pageInfo req.Get{{.StructName}}ListReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := service{{ .StructName }}.Get{{.StructName}}ListByPage(pageInfo)
	if err != nil {
		global.RLB_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		Data:     list,
		Total:    total,
		Page:     pageInfo.Current,
		PageSize: pageInfo.PageSize,
		Success:  true,
	}, "获取成功", c)
}


// Save{{.StructName}} 保存{{.Description}}
// @Tags {{.StructName}}
// @Summary 保存{{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "保存{{.Description}}"
// @Success 200 {object} response.Response{msg=string} "保存成功"
// @Router /{{.Abbreviation}}/create{{.StructName}} [post]
func ({{.Abbreviation}} *{{.Abbreviation}}Api) Save{{.StructName}}(c *gin.Context) {
	var info model.{{.StructName}}
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := service{{ .StructName }}.Save{{.StructName}}(&info); err != nil {
		global.RLB_LOG.Error("保存失败!", zap.Error(err))
		response.FailWithMessage("保存失败", c)
		return
	}
	response.OkWithMessage("保存成功", c)
}


// Delete{{.StructName}} 删除{{.Description}}
// @Tags {{.StructName}}
// @Summary 删除{{.Description}}
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "删除{{.Description}}"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [delete]
func ({{.Abbreviation}} *{{.Abbreviation}}Api) Delete{{.StructName}}(c *gin.Context) {
	id := c.Query("id")
	idInt, _ := utils.StrToInt64(id)
	resultInfo, err := service{{ .StructName }}.Delete{{.StructName}}(idInt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resultInfo.RowsAffected > 0 {
		response.OkWithMessage("删除"+fmt.Sprintf("%d", resultInfo.RowsAffected)+"条", c)

	} else {
		response.FailWithMessage("删除失败", c)
	}
}
