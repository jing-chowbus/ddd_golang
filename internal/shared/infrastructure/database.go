package infrastructure

import (
	"context"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Database interface {
	Debug() Database
	Scopes(funcs ...func(Database) Database) Database
	Unscoped() Database
	Save(value interface{}) Database
	Take(out interface{}, where ...interface{}) Database
	First(out interface{}, where ...interface{}) Database
	Find(out interface{}, where ...interface{}) Database
	Preload(column string, conditions ...interface{}) Database
	Scan(dest interface{}) Database
	Last(out interface{}, where ...interface{}) Database
	Model(value interface{}) Database
	Table(name string) Database
	Update(column string, value interface{}) Database
	Updates(values interface{}) Database
	Create(value interface{}) Database
	Select(query interface{}, args ...interface{}) Database
	Where(query interface{}, args ...interface{}) Database
	Or(query interface{}, args ...interface{}) Database
	Not(query interface{}, args ...interface{}) Database
	Raw(sql string, values ...interface{}) Database
	Exec(sql string, values ...interface{}) Database
	Delete(value interface{}, where ...interface{}) Database
	Joins(query string, args ...interface{}) Database
	Order(value interface{}) Database
	Group(fields string) Database
	Limit(limit int) Database
	Offset(offset int) Database
	Count(count *int64) Database
	Commit() Database
	Rollback() Database
	Error() error
	Begin() Database
	GormDB() *gorm.DB
	Transaction(transaction func(Database) error) error
	WithTimeout(function func(Database) Database, timeout time.Duration) Database
	Distinct(columns ...string) Database
	Pluck(column string, value interface{}) Database
	Attrs(attrs ...interface{}) Database
	Assign(attrs ...interface{}) Database
	FirstOrCreate(value interface{}, where ...interface{}) Database
}

type GormDatabaseImpl struct {
	db *gorm.DB
}

// DB return gorm db
func (g *GormDatabaseImpl) GormDB() *gorm.DB {
	return g.db
}

func (g *GormDatabaseImpl) WithTimeout(function func(Database) Database, timeout time.Duration) Database {
	timeoutCtx, cancelFunc := context.WithTimeout(context.TODO(), timeout)
	defer cancelFunc()
	return function(newGormDatabaseImplWithCtx(timeoutCtx, g.db))
}

// Begin return a new instance of Database
func (g *GormDatabaseImpl) Begin() Database {
	return g.newGormDatabaseImpl(g.db.Begin())
}

func (g *GormDatabaseImpl) Transaction(transaction func(Database) error) error {
	err := g.db.Transaction(func(tx *gorm.DB) error {
		err := transaction(g.newGormDatabaseImpl(tx))
		return err
	})
	return err
}

func (g *GormDatabaseImpl) Commit() Database {
	return g.newGormDatabaseImpl(g.db.Commit())
}

func (g *GormDatabaseImpl) Rollback() Database {
	return g.newGormDatabaseImpl(g.db.Rollback())
}

func (g *GormDatabaseImpl) Debug() Database {
	return g.newGormDatabaseImpl(g.db.Debug())
}

func (g *GormDatabaseImpl) Scopes(funcs ...func(Database) Database) Database {
	var db Database
	for _, f := range funcs {
		db = f(g).(*GormDatabaseImpl)
	}
	return db
}

func (g *GormDatabaseImpl) Unscoped() Database {
	return g.newGormDatabaseImpl(g.db.Unscoped())
}

func (g *GormDatabaseImpl) Clauses(conds ...clause.Expression) Database {
	return g.newGormDatabaseImpl(g.db.Clauses(conds...))
}

func (g *GormDatabaseImpl) Save(value interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Save(value))
}

func (g *GormDatabaseImpl) Model(value interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Model(value))
}

func (g *GormDatabaseImpl) Table(name string) Database {
	return g.newGormDatabaseImpl(g.db.Table(name))
}

func (g *GormDatabaseImpl) Update(column string, value interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Update(column, value))
}

func (g *GormDatabaseImpl) Updates(values interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Updates(values))
}

func (g *GormDatabaseImpl) Create(value interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Create(value))
}

func (g *GormDatabaseImpl) Select(query interface{}, args ...interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Select(query, args...))
}

func (g *GormDatabaseImpl) Where(query interface{}, args ...interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Where(query, args...))
}

func (g *GormDatabaseImpl) Or(query interface{}, args ...interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Or(query, args...))
}

func (g *GormDatabaseImpl) Not(query interface{}, args ...interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Not(query, args...))
}

func (g *GormDatabaseImpl) Raw(sql string, values ...interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Raw(sql, values...))
}

func (g *GormDatabaseImpl) Exec(sql string, values ...interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Exec(sql, values...))
}

func (g *GormDatabaseImpl) Joins(query string, args ...interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Joins(query, args...))
}

func (g *GormDatabaseImpl) Delete(value interface{}, where ...interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Delete(value, where...))
}

func (g *GormDatabaseImpl) Last(out interface{}, where ...interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Last(out, where...))
}

func (g *GormDatabaseImpl) Find(out interface{}, where ...interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Find(out, where...))
}

func (g *GormDatabaseImpl) Group(fields string) Database {
	return g.newGormDatabaseImpl(g.db.Group(fields))
}

func (g *GormDatabaseImpl) Preload(column string, conditions ...interface{}) Database {
	var newConditions []interface{}
	for _, condition := range conditions {
		conditionFun, ok := condition.(func(db Database) Database)
		if ok {
			condition = func(db *gorm.DB) *gorm.DB {
				database := conditionFun(g.newGormDatabaseImpl(db))
				return database.GormDB()
			}
		}
		newConditions = append(newConditions, condition)
	}
	return g.newGormDatabaseImpl(g.db.Preload(column, newConditions...))
}

func (g *GormDatabaseImpl) Scan(dest interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Scan(dest))
}

func (g *GormDatabaseImpl) Take(out interface{}, where ...interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Take(out, where...))
}

func (g *GormDatabaseImpl) First(out interface{}, where ...interface{}) Database {
	return g.newGormDatabaseImpl(g.db.First(out, where...))
}

func (g *GormDatabaseImpl) Order(value interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Order(value))
}

func (g *GormDatabaseImpl) Limit(limit int) Database {
	return g.newGormDatabaseImpl(g.db.Limit(limit))
}

func (g *GormDatabaseImpl) Offset(offset int) Database {
	return g.newGormDatabaseImpl(g.db.Offset(offset))
}

func (g *GormDatabaseImpl) Count(count *int64) Database {
	return g.newGormDatabaseImpl(g.db.Count(count))
}

func (g *GormDatabaseImpl) Distinct(columns ...string) Database {
	return g.newGormDatabaseImpl(g.db.Distinct(columns))
}

func (g *GormDatabaseImpl) Pluck(column string, value interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Pluck(column, value))
}

func (g *GormDatabaseImpl) Attrs(attrs ...interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Attrs(attrs...))
}

func (g *GormDatabaseImpl) Assign(attrs ...interface{}) Database {
	return g.newGormDatabaseImpl(g.db.Assign(attrs...))
}

func (g *GormDatabaseImpl) FirstOrCreate(value interface{}, where ...interface{}) Database {
	return g.newGormDatabaseImpl(g.db.FirstOrCreate(value, where...))
}

// Error returns the error
func (g *GormDatabaseImpl) Error() error {
	return g.db.Error
}

func NewGormDatabase() (Database, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &GormDatabaseImpl{
		db: db,
	}, nil
}

func newGormDatabaseImplWithCtx(ctx context.Context, db *gorm.DB) Database {
	return &GormDatabaseImpl{
		db: db.WithContext(ctx),
	}
}

func (g *GormDatabaseImpl) newGormDatabaseImpl(db *gorm.DB) Database {
	return &GormDatabaseImpl{
		db: db,
	}
}
