
package main


import (
	"fmt"
	"github.com/psychedelicnekopunch/go-sample/scope"
)


var scopeVar = "test"

var (
	scopeVar2 = "test2"
	scopeVar3 = 3
)

// syntax error: non-declaration statement outside function body
// scopeVar4 := true


func main() {
	// fmt.Print(scopeVar, "\n")
	// scope.Test()

	fn := scope.NewFunc()
	// vars := fn.GetFuncVars()

	var vars scope.FuncVars = fn.GetFuncVars()

	fmt.Print(vars, "\n")
}
