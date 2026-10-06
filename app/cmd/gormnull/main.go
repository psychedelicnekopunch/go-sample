package main


import (
	"fmt"
	"time"

	"github.com/psychedelicnekopunch/go-sample/internal/infrastructure"
)


/*
--
-- テーブルの構造 `books`
--
CREATE TABLE `books` (
  `id` int UNSIGNED NOT NULL,
  `title` text NOT NULL,
  `description` text,
  `publish_at` int UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- テーブルのインデックス `books`
--
ALTER TABLE `books`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `books`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT;
*/
type Books struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description *string `json:"description"` // NULL
	PublishAt *int64 `json:"publishAt"` // NULL
}


func main() {

	d := infrastructure.NewDB()
	db := d.Connect()

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
		fmt.Print(*foundBook.PublishAt, "\n")
	}
}
