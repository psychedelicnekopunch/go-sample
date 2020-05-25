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
	Connection *gorm.DB
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
	d.Connection = db
	return d
}


// Begin begins a transaction
func (db *DB) Begin() *gorm.DB {
	return db.Connection.Begin()
}


func (db *DB) Connect() *gorm.DB {
	return db.Connection
}
