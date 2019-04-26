package main


import (
	"fmt"
	"time"

	"psychedelicnekopunch/go-sample/infrastructure"
)


type Users struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	IsStudent bool `json:"isStudent"`
	CreatedAt int64 `json:"createdAt"`
}


func main() {

	db := infrastructure.NewDB()

	// SQL
	rows, err := db.Raw("show tables").Rows()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			panic(err.Error())
		}
		db.Exec("TRUNCATE TABLE test." + table)
	}


	// Create
	user := Users{
		Name: "test",
		Age:  16,
		IsStudent: true,
		CreatedAt: time.Now().Unix(),
	}
	if !db.NewRecord(&user) {
		panic("could not create new record")
	}
	if err := db.Create(&user).Error; err != nil {
		panic(err.Error())
	}


	// Query
	foundUsers := []Users{}
	db.Find(&foundUsers)
	fmt.Println(foundUsers)


	// Save
	foundUser := Users{}
	db.First(&foundUser, user.ID)
	if foundUser.ID == 0 {
		panic("user not found")
	}
	foundUser.Age = 25
	foundUser.IsStudent = false
	if err := db.Save(&foundUser).Error; err != nil {
		panic(err.Error())
	}


	// Query2
	foundUsers2 := []Users{}
	db.Where("age > ? and is_student = ?", 20, false).Find(&foundUsers2)
	fmt.Println(foundUsers2)
}