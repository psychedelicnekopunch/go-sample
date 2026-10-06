package scope

import "fmt"


// scopeVar redeclared in this block
// var scopeVar = "test"

// scopeVar2 redeclared in this block
// scopeVar3 redeclared in this block
// var (
// 	scopeVar2 = "test2"
// 	scopeVar3 = 3
// )


func testFunc2() {
	fmt.Printf("func2, ")
	testFunc3()
}

func Test2() {
	testFunc2()
}


func Test3() {
	testFunc2()
}
