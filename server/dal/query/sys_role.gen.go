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

func newSysRole(db *gorm.DB, opts ...gen.DOOption) sysRole {
	_sysRole := sysRole{}

	_sysRole.sysRoleDo.UseDB(db, opts...)
	_sysRole.sysRoleDo.UseModel(&model.SysRole{})

	tableName := _sysRole.sysRoleDo.TableName()
	_sysRole.ALL = field.NewAsterisk(tableName)
	_sysRole.ID = field.NewField(tableName, "id")
	_sysRole.UUID = field.NewString(tableName, "uuid")
	_sysRole.RoleName = field.NewString(tableName, "role_name")
	_sysRole.ParentID = field.NewString(tableName, "parent_id")
	_sysRole.DefaultRouter = field.NewString(tableName, "default_router")
	_sysRole.CreateTime = field.NewTime(tableName, "create_time")
	_sysRole.UpdateTime = field.NewTime(tableName, "update_time")
	_sysRole.DeletedAt = field.NewTime(tableName, "deleted_at")

	_sysRole.fillFieldMap()

	return _sysRole
}

// sysRole 角色表
type sysRole struct {
	sysRoleDo sysRoleDo

	ALL           field.Asterisk
	ID            field.Field // 角色ID
	UUID          field.String
	RoleName      field.String // 角色名称
	ParentID      field.String // 父角色ID
	DefaultRouter field.String // 默认路由
	CreateTime    field.Time
	UpdateTime    field.Time
	DeletedAt     field.Time

	fieldMap map[string]field.Expr
}

func (s sysRole) Table(newTableName string) *sysRole {
	s.sysRoleDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysRole) As(alias string) *sysRole {
	s.sysRoleDo.DO = *(s.sysRoleDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysRole) updateTableName(table string) *sysRole {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewField(table, "id")
	s.UUID = field.NewString(table, "uuid")
	s.RoleName = field.NewString(table, "role_name")
	s.ParentID = field.NewString(table, "parent_id")
	s.DefaultRouter = field.NewString(table, "default_router")
	s.CreateTime = field.NewTime(table, "create_time")
	s.UpdateTime = field.NewTime(table, "update_time")
	s.DeletedAt = field.NewTime(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *sysRole) WithContext(ctx context.Context) *sysRoleDo { return s.sysRoleDo.WithContext(ctx) }

func (s sysRole) TableName() string { return s.sysRoleDo.TableName() }

func (s sysRole) Alias() string { return s.sysRoleDo.Alias() }

func (s sysRole) Columns(cols ...field.Expr) gen.Columns { return s.sysRoleDo.Columns(cols...) }

func (s *sysRole) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysRole) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["id"] = s.ID
	s.fieldMap["uuid"] = s.UUID
	s.fieldMap["role_name"] = s.RoleName
	s.fieldMap["parent_id"] = s.ParentID
	s.fieldMap["default_router"] = s.DefaultRouter
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["update_time"] = s.UpdateTime
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s sysRole) clone(db *gorm.DB) sysRole {
	s.sysRoleDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysRole) replaceDB(db *gorm.DB) sysRole {
	s.sysRoleDo.ReplaceDB(db)
	return s
}

type sysRoleDo struct{ gen.DO }

func (s sysRoleDo) Debug() *sysRoleDo {
	return s.withDO(s.DO.Debug())
}

func (s sysRoleDo) WithContext(ctx context.Context) *sysRoleDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysRoleDo) ReadDB() *sysRoleDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysRoleDo) WriteDB() *sysRoleDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysRoleDo) Session(config *gorm.Session) *sysRoleDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysRoleDo) Clauses(conds ...clause.Expression) *sysRoleDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysRoleDo) Returning(value interface{}, columns ...string) *sysRoleDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysRoleDo) Not(conds ...gen.Condition) *sysRoleDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysRoleDo) Or(conds ...gen.Condition) *sysRoleDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysRoleDo) Select(conds ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysRoleDo) Where(conds ...gen.Condition) *sysRoleDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysRoleDo) Order(conds ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysRoleDo) Distinct(cols ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysRoleDo) Omit(cols ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysRoleDo) Join(table schema.Tabler, on ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysRoleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysRoleDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysRoleDo) Group(cols ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysRoleDo) Having(conds ...gen.Condition) *sysRoleDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysRoleDo) Limit(limit int) *sysRoleDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysRoleDo) Offset(offset int) *sysRoleDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysRoleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysRoleDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysRoleDo) Unscoped() *sysRoleDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysRoleDo) Create(values ...*model.SysRole) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysRoleDo) CreateInBatches(values []*model.SysRole, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysRoleDo) Save(values ...*model.SysRole) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysRoleDo) First() (*model.SysRole, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRole), nil
	}
}

func (s sysRoleDo) Take() (*model.SysRole, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRole), nil
	}
}

func (s sysRoleDo) Last() (*model.SysRole, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRole), nil
	}
}

func (s sysRoleDo) Find() ([]*model.SysRole, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysRole), err
}

func (s sysRoleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysRole, err error) {
	buf := make([]*model.SysRole, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysRoleDo) FindInBatches(result *[]*model.SysRole, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysRoleDo) Attrs(attrs ...field.AssignExpr) *sysRoleDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysRoleDo) Assign(attrs ...field.AssignExpr) *sysRoleDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysRoleDo) Joins(fields ...field.RelationField) *sysRoleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysRoleDo) Preload(fields ...field.RelationField) *sysRoleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysRoleDo) FirstOrInit() (*model.SysRole, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRole), nil
	}
}

func (s sysRoleDo) FirstOrCreate() (*model.SysRole, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRole), nil
	}
}

func (s sysRoleDo) FindByPage(offset int, limit int) (result []*model.SysRole, count int64, err error) {
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

func (s sysRoleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysRoleDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysRoleDo) Delete(models ...*model.SysRole) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysRoleDo) withDO(do gen.Dao) *sysRoleDo {
	s.DO = *do.(*gen.DO)
	return s
}
