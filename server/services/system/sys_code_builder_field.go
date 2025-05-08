package system

import (
	"context"

	"github.com/kesilent/react-light-blog/dal/model"
	"github.com/kesilent/react-light-blog/dal/query"
	req "github.com/kesilent/react-light-blog/dal/request"
)

type CodeBuilderFieldService struct{}

var CodeBuilderFieldServiceApp = new(CodeBuilderFieldService)

/**
 * @Author: Yang
 * @description: 获取分页列表
 * @param {req.GetCodeBuilderFieldReq} info
 * @return {*}
 */
func (codeBuilderFieldService *CodeBuilderFieldService) GetListByPage(info req.GetCodeBuilderFieldReq) (list []*model.SysCodeBuilderField, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Current - 1)

	db := query.Q.SysCodeBuilderField.WithContext(context.Background())

	total, err = db.Count()
	if err != nil {
		return
	}

	list, err = db.Limit(limit).Offset(offset).Find()

	return list, total, err
}
