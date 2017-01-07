package models

import "fmt"

func testFunc2() {
	fmt.Printf("func2, ")
	testFunc3()
}

func Func2() {
	testFunc2()
}


func Func3() {
	testFunc2()
}