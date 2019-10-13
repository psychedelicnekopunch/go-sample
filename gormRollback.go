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


type Books struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description *string `json:"description"`
	PublishAt *int64 `json:"publishAt"`
}


func main() {

	db := infrastructure.NewDB()

	tx := db.Begin()

	/**
	 * Create: 作成
	 */
	user := Users{
		Name: "rollback test",
		Age:  16,
		IsStudent: true,
		CreatedAt: time.Now().Unix(),
	}
	if !tx.NewRecord(&user) {
		tx.Rollback()
		panic("could not create new record")
	}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		panic(err.Error())
	}


	book := Books{
		Title: "rollback test",
	}
	if !tx.NewRecord(&book) {
		tx.Rollback()
		panic("could not create new record")
	}
	if err := tx.Create(&book).Error; err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	// Commit() しないと更新が反映されない
	// tx.Commit()
	// tx.Rollback()

	/**
	 * Query: 参照
	 */
	foundUsers := []Users{}
	// SELECT * FROM users;
	db.Find(&foundUsers)
	if len(foundUsers) == 0 {
		fmt.Printf("not found users")
	}
	fmt.Println(foundUsers)


	foundBooks := []Books{}
	// SELECT * FROM books;
	db.Find(&foundBooks)
	if len(foundBooks) == 0 {
		fmt.Printf("not found books")
	}
	fmt.Println(foundBooks)
}
