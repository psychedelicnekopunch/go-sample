
package main


import (
	"fmt"
)


type User struct {
	ID int
	Name string
}


func main() {
	// 1. var 変数 型
	// 2. var 変数 型 = 値
	// 3. var 変数 = 値
	// 4. 変数 := 値
	// 5. 変数 := 値（ここでは空のスライスを代入）

	var users1 []User
	var users2 []User = []User{ { ID: 1, Name: "sample" } }
	var users3 = []User{ { ID: 1, Name: "sample" } }
	users4 := []User{ { ID: 1, Name: "sample" } }
	users5 := []User{}

	fmt.Print(users1)// []
	fmt.Print(users2)// [{1 sample}]
	fmt.Print(users3)// [{1 sample}]
	fmt.Print(users4)// [{1 sample}]
	fmt.Print(users5)// []


	// 複数の宣言その1
	var test1, test2 int
	var test3, test4 string = "test3", "test4"
	var test5, test6 = "test5", "test6"
	test7, test8 := "test7", "test8"

	fmt.Print(test1, test2)// 0 0
	fmt.Printf(test3 + ", " + test4 + "\n")// test3, test4
	fmt.Printf(test5 + ", " + test6 + "\n")// test5, test6
	fmt.Printf(test7 + ", " + test8 + "\n")// test7, test8


	// 複数の宣言その2
	var (
		i1 int
		s1 string
	)
	var (
		i2 int = 2
		s2 string = "test2"
	)
	var (
		i3 = 3
		s3 = "test3"
	)

	fmt.Print(i1, s1)// 0, ""
	fmt.Print(i2, s2)// 2, test2
	fmt.Print(i3, s3)// 3, test3
}