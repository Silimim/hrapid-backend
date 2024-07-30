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

	"github.com/Silimim/hrapid-backend/db/model"
)

func newList(db *gorm.DB, opts ...gen.DOOption) list {
	_list := list{}

	_list.listDo.UseDB(db, opts...)
	_list.listDo.UseModel(&model.List{})

	tableName := _list.listDo.TableName()
	_list.ALL = field.NewAsterisk(tableName)
	_list.ID = field.NewInt32(tableName, "id")
	_list.Description = field.NewString(tableName, "description")
	_list.DateAdded = field.NewTime(tableName, "date_added")
	_list.UserAddedID = field.NewInt32(tableName, "user_added_id")

	_list.fillFieldMap()

	return _list
}

type list struct {
	listDo

	ALL         field.Asterisk
	ID          field.Int32
	Description field.String
	DateAdded   field.Time
	UserAddedID field.Int32

	fieldMap map[string]field.Expr
}

func (l list) Table(newTableName string) *list {
	l.listDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l list) As(alias string) *list {
	l.listDo.DO = *(l.listDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *list) updateTableName(table string) *list {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt32(table, "id")
	l.Description = field.NewString(table, "description")
	l.DateAdded = field.NewTime(table, "date_added")
	l.UserAddedID = field.NewInt32(table, "user_added_id")

	l.fillFieldMap()

	return l
}

func (l *list) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *list) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 4)
	l.fieldMap["id"] = l.ID
	l.fieldMap["description"] = l.Description
	l.fieldMap["date_added"] = l.DateAdded
	l.fieldMap["user_added_id"] = l.UserAddedID
}

func (l list) clone(db *gorm.DB) list {
	l.listDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l list) replaceDB(db *gorm.DB) list {
	l.listDo.ReplaceDB(db)
	return l
}

type listDo struct{ gen.DO }

type IListDo interface {
	gen.SubQuery
	Debug() IListDo
	WithContext(ctx context.Context) IListDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IListDo
	WriteDB() IListDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IListDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IListDo
	Not(conds ...gen.Condition) IListDo
	Or(conds ...gen.Condition) IListDo
	Select(conds ...field.Expr) IListDo
	Where(conds ...gen.Condition) IListDo
	Order(conds ...field.Expr) IListDo
	Distinct(cols ...field.Expr) IListDo
	Omit(cols ...field.Expr) IListDo
	Join(table schema.Tabler, on ...field.Expr) IListDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IListDo
	RightJoin(table schema.Tabler, on ...field.Expr) IListDo
	Group(cols ...field.Expr) IListDo
	Having(conds ...gen.Condition) IListDo
	Limit(limit int) IListDo
	Offset(offset int) IListDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IListDo
	Unscoped() IListDo
	Create(values ...*model.List) error
	CreateInBatches(values []*model.List, batchSize int) error
	Save(values ...*model.List) error
	First() (*model.List, error)
	Take() (*model.List, error)
	Last() (*model.List, error)
	Find() ([]*model.List, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.List, err error)
	FindInBatches(result *[]*model.List, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.List) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IListDo
	Assign(attrs ...field.AssignExpr) IListDo
	Joins(fields ...field.RelationField) IListDo
	Preload(fields ...field.RelationField) IListDo
	FirstOrInit() (*model.List, error)
	FirstOrCreate() (*model.List, error)
	FindByPage(offset int, limit int) (result []*model.List, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IListDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l listDo) Debug() IListDo {
	return l.withDO(l.DO.Debug())
}

func (l listDo) WithContext(ctx context.Context) IListDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l listDo) ReadDB() IListDo {
	return l.Clauses(dbresolver.Read)
}

func (l listDo) WriteDB() IListDo {
	return l.Clauses(dbresolver.Write)
}

func (l listDo) Session(config *gorm.Session) IListDo {
	return l.withDO(l.DO.Session(config))
}

func (l listDo) Clauses(conds ...clause.Expression) IListDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l listDo) Returning(value interface{}, columns ...string) IListDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l listDo) Not(conds ...gen.Condition) IListDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l listDo) Or(conds ...gen.Condition) IListDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l listDo) Select(conds ...field.Expr) IListDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l listDo) Where(conds ...gen.Condition) IListDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l listDo) Order(conds ...field.Expr) IListDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l listDo) Distinct(cols ...field.Expr) IListDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l listDo) Omit(cols ...field.Expr) IListDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l listDo) Join(table schema.Tabler, on ...field.Expr) IListDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l listDo) LeftJoin(table schema.Tabler, on ...field.Expr) IListDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l listDo) RightJoin(table schema.Tabler, on ...field.Expr) IListDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l listDo) Group(cols ...field.Expr) IListDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l listDo) Having(conds ...gen.Condition) IListDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l listDo) Limit(limit int) IListDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l listDo) Offset(offset int) IListDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l listDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IListDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l listDo) Unscoped() IListDo {
	return l.withDO(l.DO.Unscoped())
}

func (l listDo) Create(values ...*model.List) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l listDo) CreateInBatches(values []*model.List, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l listDo) Save(values ...*model.List) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l listDo) First() (*model.List, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.List), nil
	}
}

func (l listDo) Take() (*model.List, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.List), nil
	}
}

func (l listDo) Last() (*model.List, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.List), nil
	}
}

func (l listDo) Find() ([]*model.List, error) {
	result, err := l.DO.Find()
	return result.([]*model.List), err
}

func (l listDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.List, err error) {
	buf := make([]*model.List, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l listDo) FindInBatches(result *[]*model.List, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l listDo) Attrs(attrs ...field.AssignExpr) IListDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l listDo) Assign(attrs ...field.AssignExpr) IListDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l listDo) Joins(fields ...field.RelationField) IListDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l listDo) Preload(fields ...field.RelationField) IListDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l listDo) FirstOrInit() (*model.List, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.List), nil
	}
}

func (l listDo) FirstOrCreate() (*model.List, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.List), nil
	}
}

func (l listDo) FindByPage(offset int, limit int) (result []*model.List, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l listDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l listDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l listDo) Delete(models ...*model.List) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *listDo) withDO(do gen.Dao) *listDo {
	l.DO = *do.(*gen.DO)
	return l
}
