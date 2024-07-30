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

func newListCompany(db *gorm.DB, opts ...gen.DOOption) listCompany {
	_listCompany := listCompany{}

	_listCompany.listCompanyDo.UseDB(db, opts...)
	_listCompany.listCompanyDo.UseModel(&model.ListCompany{})

	tableName := _listCompany.listCompanyDo.TableName()
	_listCompany.ALL = field.NewAsterisk(tableName)
	_listCompany.ID = field.NewInt32(tableName, "id")
	_listCompany.CompanyID = field.NewInt32(tableName, "company_id")
	_listCompany.ListID = field.NewInt32(tableName, "list_id")

	_listCompany.fillFieldMap()

	return _listCompany
}

type listCompany struct {
	listCompanyDo

	ALL       field.Asterisk
	ID        field.Int32
	CompanyID field.Int32
	ListID    field.Int32

	fieldMap map[string]field.Expr
}

func (l listCompany) Table(newTableName string) *listCompany {
	l.listCompanyDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l listCompany) As(alias string) *listCompany {
	l.listCompanyDo.DO = *(l.listCompanyDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *listCompany) updateTableName(table string) *listCompany {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt32(table, "id")
	l.CompanyID = field.NewInt32(table, "company_id")
	l.ListID = field.NewInt32(table, "list_id")

	l.fillFieldMap()

	return l
}

func (l *listCompany) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *listCompany) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 3)
	l.fieldMap["id"] = l.ID
	l.fieldMap["company_id"] = l.CompanyID
	l.fieldMap["list_id"] = l.ListID
}

func (l listCompany) clone(db *gorm.DB) listCompany {
	l.listCompanyDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l listCompany) replaceDB(db *gorm.DB) listCompany {
	l.listCompanyDo.ReplaceDB(db)
	return l
}

type listCompanyDo struct{ gen.DO }

type IListCompanyDo interface {
	gen.SubQuery
	Debug() IListCompanyDo
	WithContext(ctx context.Context) IListCompanyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IListCompanyDo
	WriteDB() IListCompanyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IListCompanyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IListCompanyDo
	Not(conds ...gen.Condition) IListCompanyDo
	Or(conds ...gen.Condition) IListCompanyDo
	Select(conds ...field.Expr) IListCompanyDo
	Where(conds ...gen.Condition) IListCompanyDo
	Order(conds ...field.Expr) IListCompanyDo
	Distinct(cols ...field.Expr) IListCompanyDo
	Omit(cols ...field.Expr) IListCompanyDo
	Join(table schema.Tabler, on ...field.Expr) IListCompanyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IListCompanyDo
	RightJoin(table schema.Tabler, on ...field.Expr) IListCompanyDo
	Group(cols ...field.Expr) IListCompanyDo
	Having(conds ...gen.Condition) IListCompanyDo
	Limit(limit int) IListCompanyDo
	Offset(offset int) IListCompanyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IListCompanyDo
	Unscoped() IListCompanyDo
	Create(values ...*model.ListCompany) error
	CreateInBatches(values []*model.ListCompany, batchSize int) error
	Save(values ...*model.ListCompany) error
	First() (*model.ListCompany, error)
	Take() (*model.ListCompany, error)
	Last() (*model.ListCompany, error)
	Find() ([]*model.ListCompany, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ListCompany, err error)
	FindInBatches(result *[]*model.ListCompany, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ListCompany) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IListCompanyDo
	Assign(attrs ...field.AssignExpr) IListCompanyDo
	Joins(fields ...field.RelationField) IListCompanyDo
	Preload(fields ...field.RelationField) IListCompanyDo
	FirstOrInit() (*model.ListCompany, error)
	FirstOrCreate() (*model.ListCompany, error)
	FindByPage(offset int, limit int) (result []*model.ListCompany, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IListCompanyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l listCompanyDo) Debug() IListCompanyDo {
	return l.withDO(l.DO.Debug())
}

func (l listCompanyDo) WithContext(ctx context.Context) IListCompanyDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l listCompanyDo) ReadDB() IListCompanyDo {
	return l.Clauses(dbresolver.Read)
}

func (l listCompanyDo) WriteDB() IListCompanyDo {
	return l.Clauses(dbresolver.Write)
}

func (l listCompanyDo) Session(config *gorm.Session) IListCompanyDo {
	return l.withDO(l.DO.Session(config))
}

func (l listCompanyDo) Clauses(conds ...clause.Expression) IListCompanyDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l listCompanyDo) Returning(value interface{}, columns ...string) IListCompanyDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l listCompanyDo) Not(conds ...gen.Condition) IListCompanyDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l listCompanyDo) Or(conds ...gen.Condition) IListCompanyDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l listCompanyDo) Select(conds ...field.Expr) IListCompanyDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l listCompanyDo) Where(conds ...gen.Condition) IListCompanyDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l listCompanyDo) Order(conds ...field.Expr) IListCompanyDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l listCompanyDo) Distinct(cols ...field.Expr) IListCompanyDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l listCompanyDo) Omit(cols ...field.Expr) IListCompanyDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l listCompanyDo) Join(table schema.Tabler, on ...field.Expr) IListCompanyDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l listCompanyDo) LeftJoin(table schema.Tabler, on ...field.Expr) IListCompanyDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l listCompanyDo) RightJoin(table schema.Tabler, on ...field.Expr) IListCompanyDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l listCompanyDo) Group(cols ...field.Expr) IListCompanyDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l listCompanyDo) Having(conds ...gen.Condition) IListCompanyDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l listCompanyDo) Limit(limit int) IListCompanyDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l listCompanyDo) Offset(offset int) IListCompanyDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l listCompanyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IListCompanyDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l listCompanyDo) Unscoped() IListCompanyDo {
	return l.withDO(l.DO.Unscoped())
}

func (l listCompanyDo) Create(values ...*model.ListCompany) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l listCompanyDo) CreateInBatches(values []*model.ListCompany, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l listCompanyDo) Save(values ...*model.ListCompany) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l listCompanyDo) First() (*model.ListCompany, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCompany), nil
	}
}

func (l listCompanyDo) Take() (*model.ListCompany, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCompany), nil
	}
}

func (l listCompanyDo) Last() (*model.ListCompany, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCompany), nil
	}
}

func (l listCompanyDo) Find() ([]*model.ListCompany, error) {
	result, err := l.DO.Find()
	return result.([]*model.ListCompany), err
}

func (l listCompanyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ListCompany, err error) {
	buf := make([]*model.ListCompany, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l listCompanyDo) FindInBatches(result *[]*model.ListCompany, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l listCompanyDo) Attrs(attrs ...field.AssignExpr) IListCompanyDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l listCompanyDo) Assign(attrs ...field.AssignExpr) IListCompanyDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l listCompanyDo) Joins(fields ...field.RelationField) IListCompanyDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l listCompanyDo) Preload(fields ...field.RelationField) IListCompanyDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l listCompanyDo) FirstOrInit() (*model.ListCompany, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCompany), nil
	}
}

func (l listCompanyDo) FirstOrCreate() (*model.ListCompany, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCompany), nil
	}
}

func (l listCompanyDo) FindByPage(offset int, limit int) (result []*model.ListCompany, count int64, err error) {
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

func (l listCompanyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l listCompanyDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l listCompanyDo) Delete(models ...*model.ListCompany) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *listCompanyDo) withDO(do gen.Dao) *listCompanyDo {
	l.DO = *do.(*gen.DO)
	return l
}