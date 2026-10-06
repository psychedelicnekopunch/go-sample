
package main


import (
	"github.com/psychedelicnekopunch/go-sample/internal/structtest"
)


func main() {
	s := new(structtest.Sample)
	s.Do("main2.go", 3)
}
