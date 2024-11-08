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

func newCompanyEmployee(db *gorm.DB, opts ...gen.DOOption) companyEmployee {
	_companyEmployee := companyEmployee{}

	_companyEmployee.companyEmployeeDo.UseDB(db, opts...)
	_companyEmployee.companyEmployeeDo.UseModel(&model.CompanyEmployee{})

	tableName := _companyEmployee.companyEmployeeDo.TableName()
	_companyEmployee.ALL = field.NewAsterisk(tableName)
	_companyEmployee.ID = field.NewInt32(tableName, "id")
	_companyEmployee.EmployeeID = field.NewInt32(tableName, "employee_id")
	_companyEmployee.CompanyID = field.NewInt32(tableName, "company_id")

	_companyEmployee.fillFieldMap()

	return _companyEmployee
}

type companyEmployee struct {
	companyEmployeeDo

	ALL        field.Asterisk
	ID         field.Int32
	EmployeeID field.Int32
	CompanyID  field.Int32

	fieldMap map[string]field.Expr
}

func (c companyEmployee) Table(newTableName string) *companyEmployee {
	c.companyEmployeeDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c companyEmployee) As(alias string) *companyEmployee {
	c.companyEmployeeDo.DO = *(c.companyEmployeeDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *companyEmployee) updateTableName(table string) *companyEmployee {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt32(table, "id")
	c.EmployeeID = field.NewInt32(table, "employee_id")
	c.CompanyID = field.NewInt32(table, "company_id")

	c.fillFieldMap()

	return c
}

func (c *companyEmployee) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *companyEmployee) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 3)
	c.fieldMap["id"] = c.ID
	c.fieldMap["employee_id"] = c.EmployeeID
	c.fieldMap["company_id"] = c.CompanyID
}

func (c companyEmployee) clone(db *gorm.DB) companyEmployee {
	c.companyEmployeeDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c companyEmployee) replaceDB(db *gorm.DB) companyEmployee {
	c.companyEmployeeDo.ReplaceDB(db)
	return c
}

type companyEmployeeDo struct{ gen.DO }

type ICompanyEmployeeDo interface {
	gen.SubQuery
	Debug() ICompanyEmployeeDo
	WithContext(ctx context.Context) ICompanyEmployeeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICompanyEmployeeDo
	WriteDB() ICompanyEmployeeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICompanyEmployeeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICompanyEmployeeDo
	Not(conds ...gen.Condition) ICompanyEmployeeDo
	Or(conds ...gen.Condition) ICompanyEmployeeDo
	Select(conds ...field.Expr) ICompanyEmployeeDo
	Where(conds ...gen.Condition) ICompanyEmployeeDo
	Order(conds ...field.Expr) ICompanyEmployeeDo
	Distinct(cols ...field.Expr) ICompanyEmployeeDo
	Omit(cols ...field.Expr) ICompanyEmployeeDo
	Join(table schema.Tabler, on ...field.Expr) ICompanyEmployeeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICompanyEmployeeDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICompanyEmployeeDo
	Group(cols ...field.Expr) ICompanyEmployeeDo
	Having(conds ...gen.Condition) ICompanyEmployeeDo
	Limit(limit int) ICompanyEmployeeDo
	Offset(offset int) ICompanyEmployeeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICompanyEmployeeDo
	Unscoped() ICompanyEmployeeDo
	Create(values ...*model.CompanyEmployee) error
	CreateInBatches(values []*model.CompanyEmployee, batchSize int) error
	Save(values ...*model.CompanyEmployee) error
	First() (*model.CompanyEmployee, error)
	Take() (*model.CompanyEmployee, error)
	Last() (*model.CompanyEmployee, error)
	Find() ([]*model.CompanyEmployee, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CompanyEmployee, err error)
	FindInBatches(result *[]*model.CompanyEmployee, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CompanyEmployee) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICompanyEmployeeDo
	Assign(attrs ...field.AssignExpr) ICompanyEmployeeDo
	Joins(fields ...field.RelationField) ICompanyEmployeeDo
	Preload(fields ...field.RelationField) ICompanyEmployeeDo
	FirstOrInit() (*model.CompanyEmployee, error)
	FirstOrCreate() (*model.CompanyEmployee, error)
	FindByPage(offset int, limit int) (result []*model.CompanyEmployee, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICompanyEmployeeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c companyEmployeeDo) Debug() ICompanyEmployeeDo {
	return c.withDO(c.DO.Debug())
}

func (c companyEmployeeDo) WithContext(ctx context.Context) ICompanyEmployeeDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c companyEmployeeDo) ReadDB() ICompanyEmployeeDo {
	return c.Clauses(dbresolver.Read)
}

func (c companyEmployeeDo) WriteDB() ICompanyEmployeeDo {
	return c.Clauses(dbresolver.Write)
}

func (c companyEmployeeDo) Session(config *gorm.Session) ICompanyEmployeeDo {
	return c.withDO(c.DO.Session(config))
}

func (c companyEmployeeDo) Clauses(conds ...clause.Expression) ICompanyEmployeeDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c companyEmployeeDo) Returning(value interface{}, columns ...string) ICompanyEmployeeDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c companyEmployeeDo) Not(conds ...gen.Condition) ICompanyEmployeeDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c companyEmployeeDo) Or(conds ...gen.Condition) ICompanyEmployeeDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c companyEmployeeDo) Select(conds ...field.Expr) ICompanyEmployeeDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c companyEmployeeDo) Where(conds ...gen.Condition) ICompanyEmployeeDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c companyEmployeeDo) Order(conds ...field.Expr) ICompanyEmployeeDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c companyEmployeeDo) Distinct(cols ...field.Expr) ICompanyEmployeeDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c companyEmployeeDo) Omit(cols ...field.Expr) ICompanyEmployeeDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c companyEmployeeDo) Join(table schema.Tabler, on ...field.Expr) ICompanyEmployeeDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c companyEmployeeDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICompanyEmployeeDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c companyEmployeeDo) RightJoin(table schema.Tabler, on ...field.Expr) ICompanyEmployeeDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c companyEmployeeDo) Group(cols ...field.Expr) ICompanyEmployeeDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c companyEmployeeDo) Having(conds ...gen.Condition) ICompanyEmployeeDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c companyEmployeeDo) Limit(limit int) ICompanyEmployeeDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c companyEmployeeDo) Offset(offset int) ICompanyEmployeeDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c companyEmployeeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICompanyEmployeeDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c companyEmployeeDo) Unscoped() ICompanyEmployeeDo {
	return c.withDO(c.DO.Unscoped())
}

func (c companyEmployeeDo) Create(values ...*model.CompanyEmployee) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c companyEmployeeDo) CreateInBatches(values []*model.CompanyEmployee, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c companyEmployeeDo) Save(values ...*model.CompanyEmployee) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c companyEmployeeDo) First() (*model.CompanyEmployee, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CompanyEmployee), nil
	}
}

func (c companyEmployeeDo) Take() (*model.CompanyEmployee, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CompanyEmployee), nil
	}
}

func (c companyEmployeeDo) Last() (*model.CompanyEmployee, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CompanyEmployee), nil
	}
}

func (c companyEmployeeDo) Find() ([]*model.CompanyEmployee, error) {
	result, err := c.DO.Find()
	return result.([]*model.CompanyEmployee), err
}

func (c companyEmployeeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CompanyEmployee, err error) {
	buf := make([]*model.CompanyEmployee, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c companyEmployeeDo) FindInBatches(result *[]*model.CompanyEmployee, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c companyEmployeeDo) Attrs(attrs ...field.AssignExpr) ICompanyEmployeeDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c companyEmployeeDo) Assign(attrs ...field.AssignExpr) ICompanyEmployeeDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c companyEmployeeDo) Joins(fields ...field.RelationField) ICompanyEmployeeDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c companyEmployeeDo) Preload(fields ...field.RelationField) ICompanyEmployeeDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c companyEmployeeDo) FirstOrInit() (*model.CompanyEmployee, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CompanyEmployee), nil
	}
}

func (c companyEmployeeDo) FirstOrCreate() (*model.CompanyEmployee, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CompanyEmployee), nil
	}
}

func (c companyEmployeeDo) FindByPage(offset int, limit int) (result []*model.CompanyEmployee, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c companyEmployeeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c companyEmployeeDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c companyEmployeeDo) Delete(models ...*model.CompanyEmployee) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *companyEmployeeDo) withDO(do gen.Dao) *companyEmployeeDo {
	c.DO = *do.(*gen.DO)
	return c
}
