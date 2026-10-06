
package main


import (
	"fmt"
	"github.com/psychedelicnekopunch/go-sample/internal/importfile"
	// "github.com/psychedelicnekopunch/go-sample/internal/importfile/importfile"
	importfile2 "github.com/psychedelicnekopunch/go-sample/internal/importfile/importfile"
)


func main() {

	fmt.Print(importfile2.Test(), "\n")
	fmt.Print(importfile.Test(), "\n")

	fn := importfile.NewFunc()
	// vars := fn.GetFuncVars()
	var vars importfile.FuncVars = fn.GetFuncVars()

	fmt.Print(vars, "\n")
}
