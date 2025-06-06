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

func newJwtBlacklist(db *gorm.DB, opts ...gen.DOOption) jwtBlacklist {
	_jwtBlacklist := jwtBlacklist{}

	_jwtBlacklist.jwtBlacklistDo.UseDB(db, opts...)
	_jwtBlacklist.jwtBlacklistDo.UseModel(&model.JwtBlacklist{})

	tableName := _jwtBlacklist.jwtBlacklistDo.TableName()
	_jwtBlacklist.ALL = field.NewAsterisk(tableName)
	_jwtBlacklist.ID = field.NewField(tableName, "id")
	_jwtBlacklist.CreatedAt = field.NewTime(tableName, "created_at")
	_jwtBlacklist.UpdatedAt = field.NewTime(tableName, "updated_at")
	_jwtBlacklist.DeletedAt = field.NewField(tableName, "deleted_at")
	_jwtBlacklist.Jwt = field.NewString(tableName, "jwt")

	_jwtBlacklist.fillFieldMap()

	return _jwtBlacklist
}

type jwtBlacklist struct {
	jwtBlacklistDo jwtBlacklistDo

	ALL       field.Asterisk
	ID        field.Field
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Jwt       field.String // jwt

	fieldMap map[string]field.Expr
}

func (j jwtBlacklist) Table(newTableName string) *jwtBlacklist {
	j.jwtBlacklistDo.UseTable(newTableName)
	return j.updateTableName(newTableName)
}

func (j jwtBlacklist) As(alias string) *jwtBlacklist {
	j.jwtBlacklistDo.DO = *(j.jwtBlacklistDo.As(alias).(*gen.DO))
	return j.updateTableName(alias)
}

func (j *jwtBlacklist) updateTableName(table string) *jwtBlacklist {
	j.ALL = field.NewAsterisk(table)
	j.ID = field.NewField(table, "id")
	j.CreatedAt = field.NewTime(table, "created_at")
	j.UpdatedAt = field.NewTime(table, "updated_at")
	j.DeletedAt = field.NewField(table, "deleted_at")
	j.Jwt = field.NewString(table, "jwt")

	j.fillFieldMap()

	return j
}

func (j *jwtBlacklist) WithContext(ctx context.Context) *jwtBlacklistDo {
	return j.jwtBlacklistDo.WithContext(ctx)
}

func (j jwtBlacklist) TableName() string { return j.jwtBlacklistDo.TableName() }

func (j jwtBlacklist) Alias() string { return j.jwtBlacklistDo.Alias() }

func (j jwtBlacklist) Columns(cols ...field.Expr) gen.Columns {
	return j.jwtBlacklistDo.Columns(cols...)
}

func (j *jwtBlacklist) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := j.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (j *jwtBlacklist) fillFieldMap() {
	j.fieldMap = make(map[string]field.Expr, 5)
	j.fieldMap["id"] = j.ID
	j.fieldMap["created_at"] = j.CreatedAt
	j.fieldMap["updated_at"] = j.UpdatedAt
	j.fieldMap["deleted_at"] = j.DeletedAt
	j.fieldMap["jwt"] = j.Jwt
}

func (j jwtBlacklist) clone(db *gorm.DB) jwtBlacklist {
	j.jwtBlacklistDo.ReplaceConnPool(db.Statement.ConnPool)
	return j
}

func (j jwtBlacklist) replaceDB(db *gorm.DB) jwtBlacklist {
	j.jwtBlacklistDo.ReplaceDB(db)
	return j
}

type jwtBlacklistDo struct{ gen.DO }

func (j jwtBlacklistDo) Debug() *jwtBlacklistDo {
	return j.withDO(j.DO.Debug())
}

func (j jwtBlacklistDo) WithContext(ctx context.Context) *jwtBlacklistDo {
	return j.withDO(j.DO.WithContext(ctx))
}

func (j jwtBlacklistDo) ReadDB() *jwtBlacklistDo {
	return j.Clauses(dbresolver.Read)
}

func (j jwtBlacklistDo) WriteDB() *jwtBlacklistDo {
	return j.Clauses(dbresolver.Write)
}

func (j jwtBlacklistDo) Session(config *gorm.Session) *jwtBlacklistDo {
	return j.withDO(j.DO.Session(config))
}

func (j jwtBlacklistDo) Clauses(conds ...clause.Expression) *jwtBlacklistDo {
	return j.withDO(j.DO.Clauses(conds...))
}

func (j jwtBlacklistDo) Returning(value interface{}, columns ...string) *jwtBlacklistDo {
	return j.withDO(j.DO.Returning(value, columns...))
}

func (j jwtBlacklistDo) Not(conds ...gen.Condition) *jwtBlacklistDo {
	return j.withDO(j.DO.Not(conds...))
}

func (j jwtBlacklistDo) Or(conds ...gen.Condition) *jwtBlacklistDo {
	return j.withDO(j.DO.Or(conds...))
}

func (j jwtBlacklistDo) Select(conds ...field.Expr) *jwtBlacklistDo {
	return j.withDO(j.DO.Select(conds...))
}

func (j jwtBlacklistDo) Where(conds ...gen.Condition) *jwtBlacklistDo {
	return j.withDO(j.DO.Where(conds...))
}

func (j jwtBlacklistDo) Order(conds ...field.Expr) *jwtBlacklistDo {
	return j.withDO(j.DO.Order(conds...))
}

func (j jwtBlacklistDo) Distinct(cols ...field.Expr) *jwtBlacklistDo {
	return j.withDO(j.DO.Distinct(cols...))
}

func (j jwtBlacklistDo) Omit(cols ...field.Expr) *jwtBlacklistDo {
	return j.withDO(j.DO.Omit(cols...))
}

func (j jwtBlacklistDo) Join(table schema.Tabler, on ...field.Expr) *jwtBlacklistDo {
	return j.withDO(j.DO.Join(table, on...))
}

func (j jwtBlacklistDo) LeftJoin(table schema.Tabler, on ...field.Expr) *jwtBlacklistDo {
	return j.withDO(j.DO.LeftJoin(table, on...))
}

func (j jwtBlacklistDo) RightJoin(table schema.Tabler, on ...field.Expr) *jwtBlacklistDo {
	return j.withDO(j.DO.RightJoin(table, on...))
}

func (j jwtBlacklistDo) Group(cols ...field.Expr) *jwtBlacklistDo {
	return j.withDO(j.DO.Group(cols...))
}

func (j jwtBlacklistDo) Having(conds ...gen.Condition) *jwtBlacklistDo {
	return j.withDO(j.DO.Having(conds...))
}

func (j jwtBlacklistDo) Limit(limit int) *jwtBlacklistDo {
	return j.withDO(j.DO.Limit(limit))
}

func (j jwtBlacklistDo) Offset(offset int) *jwtBlacklistDo {
	return j.withDO(j.DO.Offset(offset))
}

func (j jwtBlacklistDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *jwtBlacklistDo {
	return j.withDO(j.DO.Scopes(funcs...))
}

func (j jwtBlacklistDo) Unscoped() *jwtBlacklistDo {
	return j.withDO(j.DO.Unscoped())
}

func (j jwtBlacklistDo) Create(values ...*model.JwtBlacklist) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Create(values)
}

func (j jwtBlacklistDo) CreateInBatches(values []*model.JwtBlacklist, batchSize int) error {
	return j.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (j jwtBlacklistDo) Save(values ...*model.JwtBlacklist) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Save(values)
}

func (j jwtBlacklistDo) First() (*model.JwtBlacklist, error) {
	if result, err := j.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.JwtBlacklist), nil
	}
}

func (j jwtBlacklistDo) Take() (*model.JwtBlacklist, error) {
	if result, err := j.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.JwtBlacklist), nil
	}
}

func (j jwtBlacklistDo) Last() (*model.JwtBlacklist, error) {
	if result, err := j.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.JwtBlacklist), nil
	}
}

func (j jwtBlacklistDo) Find() ([]*model.JwtBlacklist, error) {
	result, err := j.DO.Find()
	return result.([]*model.JwtBlacklist), err
}

func (j jwtBlacklistDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.JwtBlacklist, err error) {
	buf := make([]*model.JwtBlacklist, 0, batchSize)
	err = j.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (j jwtBlacklistDo) FindInBatches(result *[]*model.JwtBlacklist, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return j.DO.FindInBatches(result, batchSize, fc)
}

func (j jwtBlacklistDo) Attrs(attrs ...field.AssignExpr) *jwtBlacklistDo {
	return j.withDO(j.DO.Attrs(attrs...))
}

func (j jwtBlacklistDo) Assign(attrs ...field.AssignExpr) *jwtBlacklistDo {
	return j.withDO(j.DO.Assign(attrs...))
}

func (j jwtBlacklistDo) Joins(fields ...field.RelationField) *jwtBlacklistDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Joins(_f))
	}
	return &j
}

func (j jwtBlacklistDo) Preload(fields ...field.RelationField) *jwtBlacklistDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Preload(_f))
	}
	return &j
}

func (j jwtBlacklistDo) FirstOrInit() (*model.JwtBlacklist, error) {
	if result, err := j.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.JwtBlacklist), nil
	}
}

func (j jwtBlacklistDo) FirstOrCreate() (*model.JwtBlacklist, error) {
	if result, err := j.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.JwtBlacklist), nil
	}
}

func (j jwtBlacklistDo) FindByPage(offset int, limit int) (result []*model.JwtBlacklist, count int64, err error) {
	result, err = j.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = j.Offset(-1).Limit(-1).Count()
	return
}

func (j jwtBlacklistDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = j.Count()
	if err != nil {
		return
	}

	err = j.Offset(offset).Limit(limit).Scan(result)
	return
}

func (j jwtBlacklistDo) Scan(result interface{}) (err error) {
	return j.DO.Scan(result)
}

func (j jwtBlacklistDo) Delete(models ...*model.JwtBlacklist) (result gen.ResultInfo, err error) {
	return j.DO.Delete(models)
}

func (j *jwtBlacklistDo) withDO(do gen.Dao) *jwtBlacklistDo {
	j.DO = *do.(*gen.DO)
	return j
}
