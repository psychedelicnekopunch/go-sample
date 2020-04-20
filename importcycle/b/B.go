
package b


import (
	// "github.com/psychedelicnekopunch/go-sample/importcycle/a"
)


type B struct {
	// A2 a.A2
	A2 A2
}

func (b *B) Get() string {
	return b.A2.Get()
}
