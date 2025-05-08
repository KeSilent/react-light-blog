package system

import (
	response "github.com/kesilent/react-light-blog/dal/response"
	"github.com/kesilent/react-light-blog/global"
)

type AutoCodeService struct {
}

/**
 * @Author: Yang
 * @description: 获取数据库中所有的表
 * @return {*}
 */
func (autoCodeService *AutoCodeService) GetAllTableName() (tables []string, err error) {

	// 获取当前数据库名称
	var dbName string
	global.RLB_DB.Raw("SELECT DATABASE()").Scan(&dbName)

	// 查询所有表
	err = global.RLB_DB.Raw("SHOW TABLES FROM `" + dbName + "`").Scan(&tables).Error
	if err != nil {
		return nil, err
	}

	return tables, nil
}

/**
 * @Author: Yang
 * @description: 获取表中所有的字段
 * @param {string} tableName
 * @return {*}
 */
func (autoCodeService *AutoCodeService) GetFieldsByTableName(tableName string) (fields []response.AutoCodeField, err error) {

	// 查询所有表
	err = global.RLB_DB.Raw("SHOW FULL COLUMNS FROM `" + tableName + "`").Scan(&fields).Error
	if err != nil {
		return nil, err
	}

	return fields, nil
}
