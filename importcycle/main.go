
package main

import (
	"fmt"
	"github.com/psychedelicnekopunch/go-sample/importcycle/a"
)

func main() {
	a := a.NewA()
	fmt.Print(a.Get())
}
