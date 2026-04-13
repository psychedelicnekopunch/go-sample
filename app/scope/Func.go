package scope

import "fmt"


type Func struct {
	Vars FuncVars
}


type FuncVars struct {
	Test string
	Test2 string
}


var scopeVar = "test"

var (
	scopeVar2 = "test2"
	scopeVar3 = 3
)


func NewFunc() *Func {
	return &Func{
		Vars: FuncVars{
			Test: "test",
			Test2: "test2",
		},
	}
}

func testFunc() {
	fmt.Printf("func in Func.go\n")
}

func testFunc3() {
	fmt.Printf("func3 in Func.go\n")
}

func Test() {
	testFunc()
}

func (fn *Func) GetFuncVars() FuncVars {
	return fn.Vars
}
