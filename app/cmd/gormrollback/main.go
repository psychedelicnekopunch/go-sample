package main


import (
	"fmt"
	"time"

	"github.com/psychedelicnekopunch/go-sample/internal/infrastructure"
)


/*
--
-- テーブルの構造 `users`
--
CREATE TABLE `users` (
  `id` int UNSIGNED NOT NULL,
  `name` text NOT NULL,
  `age` int UNSIGNED NOT NULL,
  `is_student` tinyint UNSIGNED NOT NULL,
  `created_at` int UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- テーブルのインデックス `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `users`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT;
*/
type Users struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	IsStudent bool `json:"isStudent"`
	CreatedAt int64 `json:"createdAt"`
}


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
	Description *string `json:"description"`
	PublishAt *int64 `json:"publishAt"`
}


func main() {

	d := infrastructure.NewDB()

	db := d.Connect()
	tx := d.Begin()

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


	/**
	 * Commit() するとトランザクションが終了して更新内容が反映される
	 */
	// tx.Commit()

	/**
	 * Rollback() するとトランザクションが終了して更新内容は破棄される。
	 * NewRecord() した際の ID は進む。
	 */
	// tx.Rollback()


	/**
	 * Query: 参照
	 */
	foundUsers := []Users{}
	// SELECT * FROM users;
	// db.Find(&foundUsers)
	tx.Find(&foundUsers)
	if len(foundUsers) == 0 {
		fmt.Printf("not found users")
	}
	fmt.Println(foundUsers)


	// db と tx は違うコネクションなので、
	// tx で行なった DB 操作と db で行なった DB 操作は
	// それぞれ干渉しないので、下記のように tx と db を一緒に使用しないほうがいい。
	foundBooks := []Books{}
	// SELECT * FROM books;
	db.Find(&foundBooks)
	// tx.Find(&foundBooks)
	if len(foundBooks) == 0 {
		fmt.Printf("not found books")
	}
	fmt.Println(foundBooks)


	/**
	 * Commit() もしくは Rollback() すると
	 * トランザクションが終了するので、
	 * Commit(), Rollback() 後に、
	 * tx.Find() などをすると、
	 * 「sql: transaction has already been committed or rolled back」
	 * というエラーが起きる
	 */
	tx.Commit()
	// tx.Rollback()
}
