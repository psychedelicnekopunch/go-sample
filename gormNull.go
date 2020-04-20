package main


import (
	"fmt"
	"time"

	"github.com/psychedelicnekopunch/go-sample/infrastructure"
)


/*
CREATE TABLE `books` (
  `id` int(10) UNSIGNED NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `publish_at` int(10) UNSIGNED DEFAULT NULL
);
*/
type Books struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description *string `json:"description"` // NULL
	PublishAt *int64 `json:"publishAt"` // NULL
}


func main() {

	db := infrastructure.NewDB()

	/**
	 * Create: 作成
	 */
	// Description や PublishAt を定義しなかったら、
	// 初期値は NULL になる。
	// NULL ではないカラムの初期値については、
	// string だと空文字、 int だと 0 のようになる。
	book := Books{
		Title: "Scar Tissue",
	}
	if !db.NewRecord(&book) {
		panic("could not create new record")
	}
	if err := db.Create(&book).Error; err != nil {
		panic(err.Error())
	}


	/**
	 * Query: 参照
	 */
	foundBooks := []Books{}
	// SELECT * FROM books;
	db.Find(&foundBooks)
	if len(foundBooks) == 0 {
		fmt.Printf("not found books")
	}
	fmt.Println(foundBooks)


	/**
	 * Save: 更新
	 */
	description := "This book is the autobiography of Red Hot Chili Peppers vocalist Anthony Kiedis."
	publishAt := time.Now().Unix()

	foundBook := Books{}
	db.First(&foundBook, book.ID)
	// 取得できなかったら ID が初期値の 0 になっている。
	if foundBook.ID == 0 {
		panic("book not found")
	}
	// description と publish_at を更新する
	foundBook.Description = &description// &"description" のようにはできない。
	foundBook.PublishAt = &publishAt// &time.Now().Unix() のようにはできない。
	if err := db.Save(&foundBook).Error; err != nil {
		panic(err.Error())
	}

	fmt.Println(foundBook)

	// NULL チェックせずにデータを参照しようとして中身が NULL だったら落ちる。
	if foundBook.Description != nil {
		fmt.Printf("%s\n", *foundBook.Description)
	}
	if foundBook.PublishAt != nil {
		fmt.Print(*foundBook.PublishAt)
	}
}
