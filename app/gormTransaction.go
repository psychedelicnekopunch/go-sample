package main


import (
	"errors"
	"fmt"
	// "time"

	"github.com/jinzhu/gorm"

	"github.com/psychedelicnekopunch/go-sample/infrastructure"
)


type Users struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	IsStudent bool `json:"isStudent"`
	CreatedAt int64 `json:"createdAt"`
}


func create(db *gorm.DB, name string) error {
	user := Users{
		Name: name,
		Age:  1,
		IsStudent: true,
		CreatedAt: 0,
		// CreatedAt: time.Now().Unix(),
	}
	if !db.NewRecord(&user) {
		return errors.New("could not create new record")
	}
	return db.Create(&user).Error
}


func getList(db *gorm.DB) (users []Users) {
	users = []Users{}
	db.Find(&users)
	return users
}


func main() {

	d := infrastructure.NewDB()

	db := d.Connect()
	txA := d.Begin()
	txB := d.Begin()

	if err := create(txA, "txA 1");err != nil {
		txA.Rollback()
		fmt.Print(err.Error())
		return
	}

	if err := create(txB, "txB 1");err != nil {
		txB.Rollback()
		fmt.Print(err.Error())
		return
	}

	fmt.Println(getList(txA))
	fmt.Println(getList(txB))

	txB.Commit()

	fmt.Println(getList(txA))
	fmt.Println(getList(db))

	txA.Commit()

	fmt.Println(getList(db))

	fmt.Print("\n============\n")

	txA = d.Begin()
	txB = d.Begin()

	if err := create(txA, "txA 2");err != nil {
		txA.Rollback()
		fmt.Print(err.Error())
		return
	}

	if err := create(txB, "txB 2");err != nil {
		txB.Rollback()
		fmt.Print(err.Error())
		return
	}

	fmt.Println(getList(txA))
	fmt.Println(getList(txB))

	txB.Commit()

	fmt.Println(getList(txA))
	fmt.Println(getList(db))

	txA.Commit()

	fmt.Println(getList(db))
}
