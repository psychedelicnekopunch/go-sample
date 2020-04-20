package main


import (
	"fmt"
	"time"

	"github.com/psychedelicnekopunch/go-sample/infrastructure"
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

	/**
	 * 直接SQL文を書く
	 */
	// テーブル一覧を取得する
	// Exec() を使うとエラーになる
	// Raw() は Users などのモデルを必要としない書き方ができる
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
		// テーブルを空にする
		// 返り値が必要ない SQL 文は Exec() 使った方がわかりやすいかも
		db.Exec("TRUNCATE TABLE test." + table)
	}


	/**
	 * Create: 作成
	 */
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


	/**
	 * Query: 参照系その1
	 */
	foundUsers := []Users{}
	// SELECT * FROM users;
	db.Find(&foundUsers)
	if len(foundUsers) == 0 {
		fmt.Printf("not found users")
	}
	fmt.Println(foundUsers)


	/**
	 * Save: 更新
	 */
	foundUser := Users{}
	db.First(&foundUser, user.ID)
	// 取得できなかったら ID が初期値の 0 になっている。
	if foundUser.ID == 0 {
		panic("user not found")
	}
	// age と is_student を更新する
	foundUser.Age = 25
	foundUser.IsStudent = false
	if err := db.Save(&foundUser).Error; err != nil {
		panic(err.Error())
	}


	/**
	 * Query2: 参照系その2
	 */
	foundUsers2 := []Users{}
	// SELECT * FROM users WHERE age > 20 AND is_student = false;
	db.Where("age > ? and is_student = ?", 20, false).Find(&foundUsers2)
	if len(foundUsers2) == 0 {
		fmt.Printf("not found users")
	}
	fmt.Println(foundUsers2)
}