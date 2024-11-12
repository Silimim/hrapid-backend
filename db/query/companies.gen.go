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

func newCompany(db *gorm.DB, opts ...gen.DOOption) company {
	_company := company{}

	_company.companyDo.UseDB(db, opts...)
	_company.companyDo.UseModel(&model.Company{})

	tableName := _company.companyDo.TableName()
	_company.ALL = field.NewAsterisk(tableName)
	_company.ID = field.NewInt32(tableName, "id")
	_company.CompanyName = field.NewString(tableName, "company_name")
	_company.Address = field.NewString(tableName, "address")
	_company.City = field.NewString(tableName, "city")
	_company.Province = field.NewString(tableName, "province")
	_company.Country = field.NewString(tableName, "country")
	_company.Phone1 = field.NewString(tableName, "phone1")
	_company.Phone2 = field.NewString(tableName, "phone2")
	_company.Email1 = field.NewString(tableName, "email1")
	_company.Email2 = field.NewString(tableName, "email2")
	_company.Sales = field.NewFloat64(tableName, "sales")
	_company.Status = field.NewString(tableName, "status")
	_company.DateAdded = field.NewTime(tableName, "date_added")
	_company.UserAddedID = field.NewInt32(tableName, "user_added_id")

	_company.fillFieldMap()

	return _company
}

type company struct {
	companyDo

	ALL         field.Asterisk
	ID          field.Int32
	CompanyName field.String
	Address     field.String
	City        field.String
	Province    field.String
	Country     field.String
	Phone1      field.String
	Phone2      field.String
	Email1      field.String
	Email2      field.String
	Sales       field.Float64
	Status      field.String
	DateAdded   field.Time
	UserAddedID field.Int32

	fieldMap map[string]field.Expr
}

func (c company) Table(newTableName string) *company {
	c.companyDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c company) As(alias string) *company {
	c.companyDo.DO = *(c.companyDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *company) updateTableName(table string) *company {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt32(table, "id")
	c.CompanyName = field.NewString(table, "company_name")
	c.Address = field.NewString(table, "address")
	c.City = field.NewString(table, "city")
	c.Province = field.NewString(table, "province")
	c.Country = field.NewString(table, "country")
	c.Phone1 = field.NewString(table, "phone1")
	c.Phone2 = field.NewString(table, "phone2")
	c.Email1 = field.NewString(table, "email1")
	c.Email2 = field.NewString(table, "email2")
	c.Sales = field.NewFloat64(table, "sales")
	c.Status = field.NewString(table, "status")
	c.DateAdded = field.NewTime(table, "date_added")
	c.UserAddedID = field.NewInt32(table, "user_added_id")

	c.fillFieldMap()

	return c
}

func (c *company) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *company) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 14)
	c.fieldMap["id"] = c.ID
	c.fieldMap["company_name"] = c.CompanyName
	c.fieldMap["address"] = c.Address
	c.fieldMap["city"] = c.City
	c.fieldMap["province"] = c.Province
	c.fieldMap["country"] = c.Country
	c.fieldMap["phone1"] = c.Phone1
	c.fieldMap["phone2"] = c.Phone2
	c.fieldMap["email1"] = c.Email1
	c.fieldMap["email2"] = c.Email2
	c.fieldMap["sales"] = c.Sales
	c.fieldMap["status"] = c.Status
	c.fieldMap["date_added"] = c.DateAdded
	c.fieldMap["user_added_id"] = c.UserAddedID
}

func (c company) clone(db *gorm.DB) company {
	c.companyDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c company) replaceDB(db *gorm.DB) company {
	c.companyDo.ReplaceDB(db)
	return c
}

type companyDo struct{ gen.DO }

type ICompanyDo interface {
	gen.SubQuery
	Debug() ICompanyDo
	WithContext(ctx context.Context) ICompanyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICompanyDo
	WriteDB() ICompanyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICompanyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICompanyDo
	Not(conds ...gen.Condition) ICompanyDo
	Or(conds ...gen.Condition) ICompanyDo
	Select(conds ...field.Expr) ICompanyDo
	Where(conds ...gen.Condition) ICompanyDo
	Order(conds ...field.Expr) ICompanyDo
	Distinct(cols ...field.Expr) ICompanyDo
	Omit(cols ...field.Expr) ICompanyDo
	Join(table schema.Tabler, on ...field.Expr) ICompanyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICompanyDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICompanyDo
	Group(cols ...field.Expr) ICompanyDo
	Having(conds ...gen.Condition) ICompanyDo
	Limit(limit int) ICompanyDo
	Offset(offset int) ICompanyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICompanyDo
	Unscoped() ICompanyDo
	Create(values ...*model.Company) error
	CreateInBatches(values []*model.Company, batchSize int) error
	Save(values ...*model.Company) error
	First() (*model.Company, error)
	Take() (*model.Company, error)
	Last() (*model.Company, error)
	Find() ([]*model.Company, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Company, err error)
	FindInBatches(result *[]*model.Company, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Company) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICompanyDo
	Assign(attrs ...field.AssignExpr) ICompanyDo
	Joins(fields ...field.RelationField) ICompanyDo
	Preload(fields ...field.RelationField) ICompanyDo
	FirstOrInit() (*model.Company, error)
	FirstOrCreate() (*model.Company, error)
	FindByPage(offset int, limit int) (result []*model.Company, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICompanyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c companyDo) Debug() ICompanyDo {
	return c.withDO(c.DO.Debug())
}

func (c companyDo) WithContext(ctx context.Context) ICompanyDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c companyDo) ReadDB() ICompanyDo {
	return c.Clauses(dbresolver.Read)
}

func (c companyDo) WriteDB() ICompanyDo {
	return c.Clauses(dbresolver.Write)
}

func (c companyDo) Session(config *gorm.Session) ICompanyDo {
	return c.withDO(c.DO.Session(config))
}

func (c companyDo) Clauses(conds ...clause.Expression) ICompanyDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c companyDo) Returning(value interface{}, columns ...string) ICompanyDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c companyDo) Not(conds ...gen.Condition) ICompanyDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c companyDo) Or(conds ...gen.Condition) ICompanyDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c companyDo) Select(conds ...field.Expr) ICompanyDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c companyDo) Where(conds ...gen.Condition) ICompanyDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c companyDo) Order(conds ...field.Expr) ICompanyDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c companyDo) Distinct(cols ...field.Expr) ICompanyDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c companyDo) Omit(cols ...field.Expr) ICompanyDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c companyDo) Join(table schema.Tabler, on ...field.Expr) ICompanyDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c companyDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICompanyDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c companyDo) RightJoin(table schema.Tabler, on ...field.Expr) ICompanyDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c companyDo) Group(cols ...field.Expr) ICompanyDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c companyDo) Having(conds ...gen.Condition) ICompanyDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c companyDo) Limit(limit int) ICompanyDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c companyDo) Offset(offset int) ICompanyDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c companyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICompanyDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c companyDo) Unscoped() ICompanyDo {
	return c.withDO(c.DO.Unscoped())
}

func (c companyDo) Create(values ...*model.Company) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c companyDo) CreateInBatches(values []*model.Company, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c companyDo) Save(values ...*model.Company) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c companyDo) First() (*model.Company, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Company), nil
	}
}

func (c companyDo) Take() (*model.Company, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Company), nil
	}
}

func (c companyDo) Last() (*model.Company, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Company), nil
	}
}

func (c companyDo) Find() ([]*model.Company, error) {
	result, err := c.DO.Find()
	return result.([]*model.Company), err
}

func (c companyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Company, err error) {
	buf := make([]*model.Company, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c companyDo) FindInBatches(result *[]*model.Company, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c companyDo) Attrs(attrs ...field.AssignExpr) ICompanyDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c companyDo) Assign(attrs ...field.AssignExpr) ICompanyDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c companyDo) Joins(fields ...field.RelationField) ICompanyDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c companyDo) Preload(fields ...field.RelationField) ICompanyDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c companyDo) FirstOrInit() (*model.Company, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Company), nil
	}
}

func (c companyDo) FirstOrCreate() (*model.Company, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Company), nil
	}
}

func (c companyDo) FindByPage(offset int, limit int) (result []*model.Company, count int64, err error) {
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

func (c companyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c companyDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c companyDo) Delete(models ...*model.Company) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *companyDo) withDO(do gen.Dao) *companyDo {
	c.DO = *do.(*gen.DO)
	return c
}
