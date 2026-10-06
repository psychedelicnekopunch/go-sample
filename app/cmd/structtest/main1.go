
package main


import (
	"github.com/psychedelicnekopunch/go-sample/internal/structtest"
)


func main() {
	s := new(structtest.Sample)
	s.Do("main1.go", 5)

	s2 := new(structtest.Sample)
	s2.Do("main1.go 1", 2)

	s.Do("main1.go 2", 1)
}
