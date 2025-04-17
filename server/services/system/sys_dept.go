package system

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strconv"

	"github.com/kesilent/react-light-blog/dal/model"
	"github.com/kesilent/react-light-blog/dal/query"
	req "github.com/kesilent/react-light-blog/dal/request"
	"github.com/kesilent/react-light-blog/utils"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type DeptService struct{}

var DeptServiceApp = new(DeptService)

/**
 * @Author: Yang
 * @description: 获取部门列表
 * @param {req.GetDeptListReq} info
 * @return {*}
 */
func (deptService *DeptService) GetListByPage(info req.GetDeptListReq) (list []model.SysDept, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Current - 1)

	db := query.Q.SysDept.WithContext(context.Background())

	if info.DeptName != "" {
		db = db.Where(query.SysDept.DeptName.Like("%" + info.DeptName + "%"))
	}
	if info.Status != "" {
		status, err := strconv.ParseBool(info.Status) // 将字符串转换为布尔值
		if err != nil {
			return nil, 0, errors.New("invalid status value")
		}
		db = db.Where(query.SysDept.Status.Is(status))
	}

	deptList, err := db.Limit(limit).Offset(offset).Find()
	// 构建树结构
	treeDepts := deptService.buildTree(deptList, deptService.getLastParentId(deptList))

	deptService.sortModels(treeDepts)

	return treeDepts, int64(len(treeDepts)), err
}

/**
 * @Author: Yang
 * @description: 获取下拉菜单部门列表
 * @param {string} key
 * @return {*}
 */
func (deptService *DeptService) GetListByTreeSelect(key string) (list []model.SysDept, err error) {
	db := query.Q.SysDept.WithContext(context.Background())

	if key != "" {
		db = db.Where(query.SysDept.DeptName.Like("%" + key + "%"))
	}

	models, err := db.Debug().Find()

	if err != nil {
		return nil, err
	}

	treeModels := deptService.buildTree(models, 0)

	//补充一条默认值id为0，名称为空的节点
	treeModels = append(treeModels, model.SysDept{
		ID:       0,
		DeptName: "空",
		Parent:   "",
		Sort:     0,
		Status:   true,
	})
	deptService.sortModels(treeModels)

	return treeModels, nil
}

/**
 * @Author: Yang
 * @description: 保存部门
 * @param {model.SysDept} dept
 * @return {*}
 */
func (deptService *DeptService) SaveDept(dept model.SysDept) error {
	db := query.Q.SysDept.WithContext(context.Background())
	if dept.ID == 0 {
		_, err := db.Where(query.SysDept.DeptName.Eq(dept.DeptName)).First()
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("存在重复name，请修改name")
		}
		dept.ID = model.SnowflakeID(utils.GenID(1))
	}

	if dept.ParentID == 0 {
		dept.Parent = fmt.Sprintf("%d", dept.ID)
	} else {
		parentDept, err := db.Where(query.SysDept.ID.Eq(dept.ParentID)).First()
		if err != nil {
			return err
		}
		dept.Parent = parentDept.Parent + fmt.Sprintf("-%d", dept.ID)

	}
	return db.Save(&dept)
}

/**
 * @Author: Yang
 * @description: 删除部门
 * @param {int64} id
 * @return {*}
 */
func (deptService *DeptService) DeleteDept(id int64) (gen.ResultInfo, error) {
	db := query.Q.SysDept.WithContext(context.Background())
	resultInfo, err := db.Where(query.SysDept.ID.Eq(model.SnowflakeID(id))).Delete()
	if err != nil {
		return resultInfo, err
	}
	db.Where(query.SysDept.Parent.Like(fmt.Sprintf("%d-%%", id))).Delete()

	return resultInfo, nil
}

/**
 * @Author: Yang
 * @description:  获取最后一个父级ID
 * @param {[]*model.SysDept} models
 * @return {*}
 */
func (deptService *DeptService) getLastParentId(models []*model.SysDept) (lastParentId model.SnowflakeID) {
	if len(models) == 0 {
		return 0
	}
	lastParentId = models[0].ParentID
	for _, model := range models {
		if model.ParentID < lastParentId {
			lastParentId = model.ParentID
		}
	}

	return
}

/**
 * @Author: Yang
 * @description: 构建树结构
 * @param {[]*model.SysDept} models
 * @param {model.SnowflakeID} parentID
 * @return {*}
 */
func (deptService *DeptService) buildTree(models []*model.SysDept, parentID model.SnowflakeID) []model.SysDept {
	var tree []model.SysDept

	for _, model := range models {
		if model.ParentID == parentID {
			// 复制当前菜单节点
			node := *model
			// 递归构建子菜单
			node.Children = deptService.buildTree(models, model.ID)
			tree = append(tree, node)
		}
	}

	return tree
}

// 递归排序菜单
func (deptService *DeptService) sortModels(models []model.SysDept) {
	// 按 Sort 字段排序
	sort.Slice(models, func(i, j int) bool {
		return models[i].Sort < models[j].Sort
	})

	// 递归排序子菜单
	for i := range models {
		if len(models[i].Children) > 0 {
			deptService.sortModels(models[i].Children)
		}
	}
}
