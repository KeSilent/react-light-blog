// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/kesilent/react-light-blog/dal/model"
)

func newSysUserAuthority(db *gorm.DB, opts ...gen.DOOption) sysUserAuthority {
	_sysUserAuthority := sysUserAuthority{}

	_sysUserAuthority.sysUserAuthorityDo.UseDB(db, opts...)
	_sysUserAuthority.sysUserAuthorityDo.UseModel(&model.SysUserAuthority{})

	tableName := _sysUserAuthority.sysUserAuthorityDo.TableName()
	_sysUserAuthority.ALL = field.NewAsterisk(tableName)
	_sysUserAuthority.SysUserID = field.NewInt64(tableName, "sys_user_id")
	_sysUserAuthority.SysAuthorityID = field.NewInt64(tableName, "sys_authority_id")

	_sysUserAuthority.fillFieldMap()

	return _sysUserAuthority
}

// sysUserAuthority 用户角色关联表
type sysUserAuthority struct {
	sysUserAuthorityDo sysUserAuthorityDo

	ALL            field.Asterisk
	SysUserID      field.Int64 // 用户ID
	SysAuthorityID field.Int64 // 角色ID

	fieldMap map[string]field.Expr
}

func (s sysUserAuthority) Table(newTableName string) *sysUserAuthority {
	s.sysUserAuthorityDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysUserAuthority) As(alias string) *sysUserAuthority {
	s.sysUserAuthorityDo.DO = *(s.sysUserAuthorityDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysUserAuthority) updateTableName(table string) *sysUserAuthority {
	s.ALL = field.NewAsterisk(table)
	s.SysUserID = field.NewInt64(table, "sys_user_id")
	s.SysAuthorityID = field.NewInt64(table, "sys_authority_id")

	s.fillFieldMap()

	return s
}

func (s *sysUserAuthority) WithContext(ctx context.Context) *sysUserAuthorityDo {
	return s.sysUserAuthorityDo.WithContext(ctx)
}

func (s sysUserAuthority) TableName() string { return s.sysUserAuthorityDo.TableName() }

func (s sysUserAuthority) Alias() string { return s.sysUserAuthorityDo.Alias() }

func (s sysUserAuthority) Columns(cols ...field.Expr) gen.Columns {
	return s.sysUserAuthorityDo.Columns(cols...)
}

func (s *sysUserAuthority) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysUserAuthority) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 2)
	s.fieldMap["sys_user_id"] = s.SysUserID
	s.fieldMap["sys_authority_id"] = s.SysAuthorityID
}

func (s sysUserAuthority) clone(db *gorm.DB) sysUserAuthority {
	s.sysUserAuthorityDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysUserAuthority) replaceDB(db *gorm.DB) sysUserAuthority {
	s.sysUserAuthorityDo.ReplaceDB(db)
	return s
}

type sysUserAuthorityDo struct{ gen.DO }

func (s sysUserAuthorityDo) Debug() *sysUserAuthorityDo {
	return s.withDO(s.DO.Debug())
}

func (s sysUserAuthorityDo) WithContext(ctx context.Context) *sysUserAuthorityDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysUserAuthorityDo) ReadDB() *sysUserAuthorityDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysUserAuthorityDo) WriteDB() *sysUserAuthorityDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysUserAuthorityDo) Session(config *gorm.Session) *sysUserAuthorityDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysUserAuthorityDo) Clauses(conds ...clause.Expression) *sysUserAuthorityDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysUserAuthorityDo) Returning(value interface{}, columns ...string) *sysUserAuthorityDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysUserAuthorityDo) Not(conds ...gen.Condition) *sysUserAuthorityDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysUserAuthorityDo) Or(conds ...gen.Condition) *sysUserAuthorityDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysUserAuthorityDo) Select(conds ...field.Expr) *sysUserAuthorityDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysUserAuthorityDo) Where(conds ...gen.Condition) *sysUserAuthorityDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysUserAuthorityDo) Order(conds ...field.Expr) *sysUserAuthorityDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysUserAuthorityDo) Distinct(cols ...field.Expr) *sysUserAuthorityDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysUserAuthorityDo) Omit(cols ...field.Expr) *sysUserAuthorityDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysUserAuthorityDo) Join(table schema.Tabler, on ...field.Expr) *sysUserAuthorityDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysUserAuthorityDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysUserAuthorityDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysUserAuthorityDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysUserAuthorityDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysUserAuthorityDo) Group(cols ...field.Expr) *sysUserAuthorityDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysUserAuthorityDo) Having(conds ...gen.Condition) *sysUserAuthorityDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysUserAuthorityDo) Limit(limit int) *sysUserAuthorityDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysUserAuthorityDo) Offset(offset int) *sysUserAuthorityDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysUserAuthorityDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysUserAuthorityDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysUserAuthorityDo) Unscoped() *sysUserAuthorityDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysUserAuthorityDo) Create(values ...*model.SysUserAuthority) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysUserAuthorityDo) CreateInBatches(values []*model.SysUserAuthority, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysUserAuthorityDo) Save(values ...*model.SysUserAuthority) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysUserAuthorityDo) First() (*model.SysUserAuthority, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUserAuthority), nil
	}
}

func (s sysUserAuthorityDo) Take() (*model.SysUserAuthority, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUserAuthority), nil
	}
}

func (s sysUserAuthorityDo) Last() (*model.SysUserAuthority, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUserAuthority), nil
	}
}

func (s sysUserAuthorityDo) Find() ([]*model.SysUserAuthority, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysUserAuthority), err
}

func (s sysUserAuthorityDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysUserAuthority, err error) {
	buf := make([]*model.SysUserAuthority, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysUserAuthorityDo) FindInBatches(result *[]*model.SysUserAuthority, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysUserAuthorityDo) Attrs(attrs ...field.AssignExpr) *sysUserAuthorityDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysUserAuthorityDo) Assign(attrs ...field.AssignExpr) *sysUserAuthorityDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysUserAuthorityDo) Joins(fields ...field.RelationField) *sysUserAuthorityDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysUserAuthorityDo) Preload(fields ...field.RelationField) *sysUserAuthorityDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysUserAuthorityDo) FirstOrInit() (*model.SysUserAuthority, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUserAuthority), nil
	}
}

func (s sysUserAuthorityDo) FirstOrCreate() (*model.SysUserAuthority, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUserAuthority), nil
	}
}

func (s sysUserAuthorityDo) FindByPage(offset int, limit int) (result []*model.SysUserAuthority, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysUserAuthorityDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysUserAuthorityDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysUserAuthorityDo) Delete(models ...*model.SysUserAuthority) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysUserAuthorityDo) withDO(do gen.Dao) *sysUserAuthorityDo {
	s.DO = *do.(*gen.DO)
	return s
}
