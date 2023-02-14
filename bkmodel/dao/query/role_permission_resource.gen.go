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

	"tapi/bkmodel/dao/model"
)

func newRolePermissionResource(db *gorm.DB, opts ...gen.DOOption) rolePermissionResource {
	_rolePermissionResource := rolePermissionResource{}

	_rolePermissionResource.rolePermissionResourceDo.UseDB(db, opts...)
	_rolePermissionResource.rolePermissionResourceDo.UseModel(&model.RolePermissionResource{})

	tableName := _rolePermissionResource.rolePermissionResourceDo.TableName()
	_rolePermissionResource.ALL = field.NewAsterisk(tableName)
	_rolePermissionResource.ID = field.NewInt64(tableName, "id")
	_rolePermissionResource.RoleID = field.NewInt64(tableName, "role_id")
	_rolePermissionResource.Prid = field.NewInt64(tableName, "prid")
	_rolePermissionResource.Status = field.NewInt32(tableName, "status")
	_rolePermissionResource.Ctime = field.NewInt32(tableName, "ctime")
	_rolePermissionResource.Utime = field.NewInt32(tableName, "utime")

	_rolePermissionResource.fillFieldMap()

	return _rolePermissionResource
}

type rolePermissionResource struct {
	rolePermissionResourceDo rolePermissionResourceDo

	ALL    field.Asterisk
	ID     field.Int64 // 角色和权限资源关联表主键
	RoleID field.Int64 // 角色id
	Prid   field.Int64 // 权限资源id
	Status field.Int32 // 1.有效，2.无效
	Ctime  field.Int32 // 创建时间
	Utime  field.Int32 // 更新时间

	fieldMap map[string]field.Expr
}

func (r rolePermissionResource) Table(newTableName string) *rolePermissionResource {
	r.rolePermissionResourceDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r rolePermissionResource) As(alias string) *rolePermissionResource {
	r.rolePermissionResourceDo.DO = *(r.rolePermissionResourceDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *rolePermissionResource) updateTableName(table string) *rolePermissionResource {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.RoleID = field.NewInt64(table, "role_id")
	r.Prid = field.NewInt64(table, "prid")
	r.Status = field.NewInt32(table, "status")
	r.Ctime = field.NewInt32(table, "ctime")
	r.Utime = field.NewInt32(table, "utime")

	r.fillFieldMap()

	return r
}

func (r *rolePermissionResource) WithContext(ctx context.Context) *rolePermissionResourceDo {
	return r.rolePermissionResourceDo.WithContext(ctx)
}

func (r rolePermissionResource) TableName() string { return r.rolePermissionResourceDo.TableName() }

func (r rolePermissionResource) Alias() string { return r.rolePermissionResourceDo.Alias() }

func (r *rolePermissionResource) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *rolePermissionResource) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 6)
	r.fieldMap["id"] = r.ID
	r.fieldMap["role_id"] = r.RoleID
	r.fieldMap["prid"] = r.Prid
	r.fieldMap["status"] = r.Status
	r.fieldMap["ctime"] = r.Ctime
	r.fieldMap["utime"] = r.Utime
}

func (r rolePermissionResource) clone(db *gorm.DB) rolePermissionResource {
	r.rolePermissionResourceDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r rolePermissionResource) replaceDB(db *gorm.DB) rolePermissionResource {
	r.rolePermissionResourceDo.ReplaceDB(db)
	return r
}

type rolePermissionResourceDo struct{ gen.DO }

func (r rolePermissionResourceDo) Debug() *rolePermissionResourceDo {
	return r.withDO(r.DO.Debug())
}

func (r rolePermissionResourceDo) WithContext(ctx context.Context) *rolePermissionResourceDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r rolePermissionResourceDo) ReadDB() *rolePermissionResourceDo {
	return r.Clauses(dbresolver.Read)
}

func (r rolePermissionResourceDo) WriteDB() *rolePermissionResourceDo {
	return r.Clauses(dbresolver.Write)
}

func (r rolePermissionResourceDo) Session(config *gorm.Session) *rolePermissionResourceDo {
	return r.withDO(r.DO.Session(config))
}

func (r rolePermissionResourceDo) Clauses(conds ...clause.Expression) *rolePermissionResourceDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r rolePermissionResourceDo) Returning(value interface{}, columns ...string) *rolePermissionResourceDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r rolePermissionResourceDo) Not(conds ...gen.Condition) *rolePermissionResourceDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r rolePermissionResourceDo) Or(conds ...gen.Condition) *rolePermissionResourceDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r rolePermissionResourceDo) Select(conds ...field.Expr) *rolePermissionResourceDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r rolePermissionResourceDo) Where(conds ...gen.Condition) *rolePermissionResourceDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r rolePermissionResourceDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *rolePermissionResourceDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r rolePermissionResourceDo) Order(conds ...field.Expr) *rolePermissionResourceDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r rolePermissionResourceDo) Distinct(cols ...field.Expr) *rolePermissionResourceDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r rolePermissionResourceDo) Omit(cols ...field.Expr) *rolePermissionResourceDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r rolePermissionResourceDo) Join(table schema.Tabler, on ...field.Expr) *rolePermissionResourceDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r rolePermissionResourceDo) LeftJoin(table schema.Tabler, on ...field.Expr) *rolePermissionResourceDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r rolePermissionResourceDo) RightJoin(table schema.Tabler, on ...field.Expr) *rolePermissionResourceDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r rolePermissionResourceDo) Group(cols ...field.Expr) *rolePermissionResourceDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r rolePermissionResourceDo) Having(conds ...gen.Condition) *rolePermissionResourceDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r rolePermissionResourceDo) Limit(limit int) *rolePermissionResourceDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r rolePermissionResourceDo) Offset(offset int) *rolePermissionResourceDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r rolePermissionResourceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *rolePermissionResourceDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r rolePermissionResourceDo) Unscoped() *rolePermissionResourceDo {
	return r.withDO(r.DO.Unscoped())
}

func (r rolePermissionResourceDo) Create(values ...*model.RolePermissionResource) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r rolePermissionResourceDo) CreateInBatches(values []*model.RolePermissionResource, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r rolePermissionResourceDo) Save(values ...*model.RolePermissionResource) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r rolePermissionResourceDo) First() (*model.RolePermissionResource, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RolePermissionResource), nil
	}
}

func (r rolePermissionResourceDo) Take() (*model.RolePermissionResource, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RolePermissionResource), nil
	}
}

func (r rolePermissionResourceDo) Last() (*model.RolePermissionResource, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RolePermissionResource), nil
	}
}

func (r rolePermissionResourceDo) Find() ([]*model.RolePermissionResource, error) {
	result, err := r.DO.Find()
	return result.([]*model.RolePermissionResource), err
}

func (r rolePermissionResourceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RolePermissionResource, err error) {
	buf := make([]*model.RolePermissionResource, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r rolePermissionResourceDo) FindInBatches(result *[]*model.RolePermissionResource, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r rolePermissionResourceDo) Attrs(attrs ...field.AssignExpr) *rolePermissionResourceDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r rolePermissionResourceDo) Assign(attrs ...field.AssignExpr) *rolePermissionResourceDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r rolePermissionResourceDo) Joins(fields ...field.RelationField) *rolePermissionResourceDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r rolePermissionResourceDo) Preload(fields ...field.RelationField) *rolePermissionResourceDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r rolePermissionResourceDo) FirstOrInit() (*model.RolePermissionResource, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RolePermissionResource), nil
	}
}

func (r rolePermissionResourceDo) FirstOrCreate() (*model.RolePermissionResource, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RolePermissionResource), nil
	}
}

func (r rolePermissionResourceDo) FindByPage(offset int, limit int) (result []*model.RolePermissionResource, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r rolePermissionResourceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r rolePermissionResourceDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r rolePermissionResourceDo) Delete(models ...*model.RolePermissionResource) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *rolePermissionResourceDo) withDO(do gen.Dao) *rolePermissionResourceDo {
	r.DO = *do.(*gen.DO)
	return r
}