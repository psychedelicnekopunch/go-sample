
package main

import (
	"fmt"
	"github.com/psychedelicnekopunch/go-sample/internal/importcycle/a"
)

func main() {
	a := a.NewA()
	fmt.Print(a.Get(), "\n")
}
