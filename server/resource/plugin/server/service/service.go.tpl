package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"{{.Module}}/plugin/{{.PackageName}}/dal/model"
	"{{.Module}}/plugin/{{.PackageName}}/dal/query"
	req "{{.Module}}/plugin/{{.PackageName}}/request"
	"{{.Module}}/utils"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var {{.StructName}}Service = new({{.Abbreviation}}Service)

type {{.Abbreviation}}Service struct {}


// Get{{.StructName}}InfoList 分页获取{{.Description}}记录
// @param req.Get{{.StructName}}ListReq info
// @return list []model.{{.StructName}}, total int64, err error
func ({{.StructName}} *{{.Abbreviation}}Service) Get{{.StructName}}ListByPage(info req.Get{{.StructName}}ListReq) (list []model.{{.StructName}}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Current - 1)

	db := query.Q.{{.StructName}}.WithContext(context.Background())

	total, err = db.Count()
	if err != nil {
		return
	}

	list, err := db.Limit(limit).Offset(offset).Find()

	return list, total, err
}


// Save{{.StructName}} 保存{{.Description}}
// @param model.{{.StructName}} model
// @return error
func ({{.StructName}} *{{.Abbreviation}}Service) Save{{.StructName}}(info *model.{{.StructName}}) error {
	db := query.Q.{{.StructName}}.WithContext(context.Background())
	if info.ID == 0 {
		info.ID = model.SnowflakeID(utils.GenID(1))
	}
	return db.Save(info)
}

// Delete{{.StructName}} 删除{{.Description}}
// @param int64 Id
// @return gen.ResultInfo, error
func ({{.StructName}} *{{.Abbreviation}}Service) Delete{{.StructName}}(id int64) (gen.ResultInfo, error) {
	q := query.Q.{{.StructName}}
	return q.WithContext(context.Background()).Where(q.ID.Eq(model.SnowflakeID(uuid))).Delete()
}