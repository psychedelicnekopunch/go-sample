
package main

import (
	"fmt"
	"psychedelicnekopunch/go-sample/importcycle/a"
)

func main() {
	a := a.NewA()
	fmt.Print(a.Get())
}
