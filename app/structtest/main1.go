
package main


import (
	"github.com/psychedelicnekopunch/go-sample/structtest/test"
)


func main() {
	s := new(test.Sample)
	s.Do("main1.go", 5)

	s2 := new(test.Sample)
	s2.Do("main1.go 1", 2)

	s.Do("main1.go 2", 1)
}
