
package infrastructure


import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/go-sql-driver/mysql" ← こっちでも動く
)


type DB struct {
	Host string
	Username string
	Password string
	DBName string
	Connect *gorm.DB
}


func NewDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host: c.DB.Production.Host,
		Username: c.DB.Production.Username,
		Password: c.DB.Production.Password,
		DBName: c.DB.Production.DBName,
	})
}

/**
 * func Open(dialect string, args ...interface{}) (db *DB, err error)
 * https://godoc.org/github.com/jinzhu/gorm#Open
 */
func newDB(d *DB) *DB {
	//
	// ex) MySQL
	// https://github.com/go-sql-driver/mysql#examples
	//
	// ex) MySQL Parameters
	// https://github.com/go-sql-driver/mysql#parameters
	db, err := gorm.Open("mysql", d.Username + ":" + d.Password + "@tcp(" + d.Host + ")/" + d.DBName + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	d.Connect = db
	return d
}


func (db *DB) Create(value interface{}) *gorm.DB {
	return db.Connect.Create(value)
}

func (db *DB) Exec(sql string, values ...interface{}) *gorm.DB {
	return db.Connect.Exec(sql, values...)
}

func (db *DB) Find(out interface{}, where ...interface{}) *gorm.DB {
	return db.Connect.Find(out, where...)
}

func (db *DB) First(out interface{}, where ...interface{}) *gorm.DB {
	return db.Connect.First(out, where...)
}

func (db *DB) NewRecord(value interface{}) bool {
	return db.Connect.NewRecord(value)
}

func (db *DB) Raw(sql string, values ...interface{}) *gorm.DB {
	return db.Connect.Raw(sql, values...)
}

func (db *DB) Save(value interface{}) *gorm.DB {
	return db.Connect.Save(value)
}

func (db *DB) Where(query interface{}, args ...interface{}) *gorm.DB {
	return db.Connect.Where(query, args...)
}

func (db *DB) Begin() *gorm.DB {
	return db.Connect.Begin()
}

func (db *DB) Commit() *gorm.DB {
	return db.Connect.Commit()
}

func (db *DB) Rollback() *gorm.DB {
	return db.Connect.Rollback()
}
