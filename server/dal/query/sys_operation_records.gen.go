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

func newSysOperationRecord(db *gorm.DB, opts ...gen.DOOption) sysOperationRecord {
	_sysOperationRecord := sysOperationRecord{}

	_sysOperationRecord.sysOperationRecordDo.UseDB(db, opts...)
	_sysOperationRecord.sysOperationRecordDo.UseModel(&model.SysOperationRecord{})

	tableName := _sysOperationRecord.sysOperationRecordDo.TableName()
	_sysOperationRecord.ALL = field.NewAsterisk(tableName)
	_sysOperationRecord.ID = field.NewField(tableName, "id")
	_sysOperationRecord.CreatedAt = field.NewTime(tableName, "created_at")
	_sysOperationRecord.UpdatedAt = field.NewTime(tableName, "updated_at")
	_sysOperationRecord.DeletedAt = field.NewField(tableName, "deleted_at")
	_sysOperationRecord.IP = field.NewString(tableName, "ip")
	_sysOperationRecord.Method = field.NewString(tableName, "method")
	_sysOperationRecord.Path = field.NewString(tableName, "path")
	_sysOperationRecord.Status = field.NewInt32(tableName, "status")
	_sysOperationRecord.Latency = field.NewString(tableName, "latency")
	_sysOperationRecord.Agent = field.NewString(tableName, "agent")
	_sysOperationRecord.ErrorMessage = field.NewString(tableName, "error_message")
	_sysOperationRecord.Body = field.NewString(tableName, "body")
	_sysOperationRecord.Resp = field.NewString(tableName, "resp")
	_sysOperationRecord.UserID = field.NewField(tableName, "user_id")

	_sysOperationRecord.fillFieldMap()

	return _sysOperationRecord
}

// sysOperationRecord 操作记录表
type sysOperationRecord struct {
	sysOperationRecordDo sysOperationRecordDo

	ALL          field.Asterisk
	ID           field.Field
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	IP           field.String // 请求ip
	Method       field.String // 请求方法
	Path         field.String // 请求路径
	Status       field.Int32  // 请求状态
	Latency      field.String // 延迟
	Agent        field.String // 代理
	ErrorMessage field.String // 错误信息
	Body         field.String // 请求Body
	Resp         field.String // 响应Body
	UserID       field.Field  // 用户id

	fieldMap map[string]field.Expr
}

func (s sysOperationRecord) Table(newTableName string) *sysOperationRecord {
	s.sysOperationRecordDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysOperationRecord) As(alias string) *sysOperationRecord {
	s.sysOperationRecordDo.DO = *(s.sysOperationRecordDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysOperationRecord) updateTableName(table string) *sysOperationRecord {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewField(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.IP = field.NewString(table, "ip")
	s.Method = field.NewString(table, "method")
	s.Path = field.NewString(table, "path")
	s.Status = field.NewInt32(table, "status")
	s.Latency = field.NewString(table, "latency")
	s.Agent = field.NewString(table, "agent")
	s.ErrorMessage = field.NewString(table, "error_message")
	s.Body = field.NewString(table, "body")
	s.Resp = field.NewString(table, "resp")
	s.UserID = field.NewField(table, "user_id")

	s.fillFieldMap()

	return s
}

func (s *sysOperationRecord) WithContext(ctx context.Context) *sysOperationRecordDo {
	return s.sysOperationRecordDo.WithContext(ctx)
}

func (s sysOperationRecord) TableName() string { return s.sysOperationRecordDo.TableName() }

func (s sysOperationRecord) Alias() string { return s.sysOperationRecordDo.Alias() }

func (s sysOperationRecord) Columns(cols ...field.Expr) gen.Columns {
	return s.sysOperationRecordDo.Columns(cols...)
}

func (s *sysOperationRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysOperationRecord) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 14)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["ip"] = s.IP
	s.fieldMap["method"] = s.Method
	s.fieldMap["path"] = s.Path
	s.fieldMap["status"] = s.Status
	s.fieldMap["latency"] = s.Latency
	s.fieldMap["agent"] = s.Agent
	s.fieldMap["error_message"] = s.ErrorMessage
	s.fieldMap["body"] = s.Body
	s.fieldMap["resp"] = s.Resp
	s.fieldMap["user_id"] = s.UserID
}

func (s sysOperationRecord) clone(db *gorm.DB) sysOperationRecord {
	s.sysOperationRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysOperationRecord) replaceDB(db *gorm.DB) sysOperationRecord {
	s.sysOperationRecordDo.ReplaceDB(db)
	return s
}

type sysOperationRecordDo struct{ gen.DO }

func (s sysOperationRecordDo) Debug() *sysOperationRecordDo {
	return s.withDO(s.DO.Debug())
}

func (s sysOperationRecordDo) WithContext(ctx context.Context) *sysOperationRecordDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysOperationRecordDo) ReadDB() *sysOperationRecordDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysOperationRecordDo) WriteDB() *sysOperationRecordDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysOperationRecordDo) Session(config *gorm.Session) *sysOperationRecordDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysOperationRecordDo) Clauses(conds ...clause.Expression) *sysOperationRecordDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysOperationRecordDo) Returning(value interface{}, columns ...string) *sysOperationRecordDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysOperationRecordDo) Not(conds ...gen.Condition) *sysOperationRecordDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysOperationRecordDo) Or(conds ...gen.Condition) *sysOperationRecordDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysOperationRecordDo) Select(conds ...field.Expr) *sysOperationRecordDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysOperationRecordDo) Where(conds ...gen.Condition) *sysOperationRecordDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysOperationRecordDo) Order(conds ...field.Expr) *sysOperationRecordDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysOperationRecordDo) Distinct(cols ...field.Expr) *sysOperationRecordDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysOperationRecordDo) Omit(cols ...field.Expr) *sysOperationRecordDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysOperationRecordDo) Join(table schema.Tabler, on ...field.Expr) *sysOperationRecordDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysOperationRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysOperationRecordDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysOperationRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysOperationRecordDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysOperationRecordDo) Group(cols ...field.Expr) *sysOperationRecordDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysOperationRecordDo) Having(conds ...gen.Condition) *sysOperationRecordDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysOperationRecordDo) Limit(limit int) *sysOperationRecordDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysOperationRecordDo) Offset(offset int) *sysOperationRecordDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysOperationRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysOperationRecordDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysOperationRecordDo) Unscoped() *sysOperationRecordDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysOperationRecordDo) Create(values ...*model.SysOperationRecord) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysOperationRecordDo) CreateInBatches(values []*model.SysOperationRecord, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysOperationRecordDo) Save(values ...*model.SysOperationRecord) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysOperationRecordDo) First() (*model.SysOperationRecord, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperationRecord), nil
	}
}

func (s sysOperationRecordDo) Take() (*model.SysOperationRecord, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperationRecord), nil
	}
}

func (s sysOperationRecordDo) Last() (*model.SysOperationRecord, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperationRecord), nil
	}
}

func (s sysOperationRecordDo) Find() ([]*model.SysOperationRecord, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysOperationRecord), err
}

func (s sysOperationRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysOperationRecord, err error) {
	buf := make([]*model.SysOperationRecord, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysOperationRecordDo) FindInBatches(result *[]*model.SysOperationRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysOperationRecordDo) Attrs(attrs ...field.AssignExpr) *sysOperationRecordDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysOperationRecordDo) Assign(attrs ...field.AssignExpr) *sysOperationRecordDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysOperationRecordDo) Joins(fields ...field.RelationField) *sysOperationRecordDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysOperationRecordDo) Preload(fields ...field.RelationField) *sysOperationRecordDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysOperationRecordDo) FirstOrInit() (*model.SysOperationRecord, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperationRecord), nil
	}
}

func (s sysOperationRecordDo) FirstOrCreate() (*model.SysOperationRecord, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperationRecord), nil
	}
}

func (s sysOperationRecordDo) FindByPage(offset int, limit int) (result []*model.SysOperationRecord, count int64, err error) {
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

func (s sysOperationRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysOperationRecordDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysOperationRecordDo) Delete(models ...*model.SysOperationRecord) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysOperationRecordDo) withDO(do gen.Dao) *sysOperationRecordDo {
	s.DO = *do.(*gen.DO)
	return s
}
