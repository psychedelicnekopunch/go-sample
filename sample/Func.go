package sample

import "fmt"

func testFunc() {
	fmt.Printf("func in Func.go\n")
}

func testFunc3() {
	fmt.Printf("func3 in Func.go\n")
}

func Func() {
	testFunc()
}