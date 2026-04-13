
package main


import (
	"fmt"
	"github.com/psychedelicnekopunch/go-sample/importFile"
	// "github.com/psychedelicnekopunch/go-sample/importFile/importFile"
	importFile2 "github.com/psychedelicnekopunch/go-sample/importFile/importFile"
)


func main() {

	fmt.Print(importFile2.Test(), "\n")
	fmt.Print(importFile.Test(), "\n")

	fn := importFile.NewFunc()
	// vars := fn.GetFuncVars()
	var vars importFile.FuncVars = fn.GetFuncVars()

	fmt.Print(vars, "\n")
}
