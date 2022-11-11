
package main


import (
	"github.com/psychedelicnekopunch/go-sample/structtest/test"
)


func main() {
	s := new(test.Sample)
	s.Do("main2.go", 3)
}
