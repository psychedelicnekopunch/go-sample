
package main


import (
	"fmt"
	"math/rand"// https://golang.org/pkg/math/rand/
	"time"
	"unicode/utf8"
)


type Random struct {
	MaxRange int
	MinRange int
	Value string
}


func NewRandom() *Random {
	r := new(Random)
	// r.MaxRange = 50
	// r.MinRange = 40
	r.MaxRange = 8
	r.MinRange = 8
	// r.Value = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r.Value = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	return r
}


func (r *Random) Get() string {

	var source rand.Source
	source = rand.NewSource(time.Now().UnixNano())

	var str []rune
	var cnt int
	var res string

	str = []rune(r.Value)
	cnt = utf8.RuneCountInString(r.Value)

	for i := 0; i < r.MaxRange; i++ {
		if r.MinRange <= i {
			if rand.New(source).Intn(r.MaxRange - r.MinRange) == 0 {
				break
			}
		}
		key := rand.New(source).Intn(cnt)
		res = res + fmt.Sprintf("%c", str[key])
	}
	return res
}


func main() {
	random := NewRandom()
	fmt.Print(random.Get())
}
