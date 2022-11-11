
package test


import (
	"fmt"
	"time"
)


type Sample struct {
	Value string
}


func (s *Sample) Do(v string, secondForSleep time.Duration) {
	fmt.Printf("\n=== start(%s) ===\n", s.Value)
	s.Value = v
	fmt.Printf("s.Value = v .... s.Value is %s \n", s.Value)
	time.Sleep(time.Second * secondForSleep)
	fmt.Printf("%s", s.Value)
	fmt.Printf("\n=== end(%s) ===\n", s.Value)
}
