package sample2

import "fmt"

func testFunc() {
	fmt.Printf("func in Func3.go\n")
}

func FuncSample() {
	fmt.Printf("func smaple in Func3.go\n")
}

func Func3() {
	testFunc()
}