package system

import (
	"context"

	"github.com/kesilent/react-light-blog/dal/model"
	"github.com/kesilent/react-light-blog/dal/query"
	req "github.com/kesilent/react-light-blog/dal/request"
)

type DeptService struct{}

var DeptServiceApp = new(DeptService)

/**
 * @Author: Yang
 * @description: 获取部门列表
 * @param {req.GetMenuListReq} info
 * @return {*}
 */
func (deptService *DeptService) GetListByPage(info req.GetDeptListReq) (list []*model.SysDept, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Current - 1)

	db := query.Q.SysDept.WithContext(context.Background())

	total, err = db.Count()
	if err != nil {
		return
	}

	deptList, err := db.Limit(limit).Offset(offset).Find()

	return deptList, int64(len(deptList)), err
}

func (deptService *DeptService) getLastParentId(menus []*model.SysDept) (lastParentId model.SnowflakeID) {
	if len(menus) == 0 {
		return 0
	}
	lastParentId = menus[0].ParentID
	for _, menu := range menus {
		if menu.ParentID < lastParentId {
			lastParentId = menu.ParentID
		}
	}

	return
}
