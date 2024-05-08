
package main


import (
	"fmt"
)


func main() {
	// m := map[int]int{}
	// m := make(map[int]int)
	m := make(map[int]int, 3)
	m[0] = 1
	m[1] = 2
	m[2] = 3

	change(m)
	fmt.Print("m is " , m, "\n")

	change2(m)
	fmt.Print("m is " , m, "\n")

	copy := m
	change(copy)
	fmt.Print("m is " , m, "\n")

	fmt.Print("-----\n")


	for key, v := range m {
		m[key] = 5
		// マップの要素のアドレスが得られない理由の一つは、
		// マップが大きくなる際に既存の要素が再びハッシングされて新たなメモリ位置へ移動するかもしれず、
		// アドレスが無効になる可能性があるからです。
		// （『プログラミング言語Go』4.3章）

		// change3(&v)
		// change3(&m[key])// invalid operation: cannot take address
		fmt.Print("for: key " , key, ", v " , v, "\n")
	}
	fmt.Print("for " , m, "\n")
}


func change(v map[int]int) {
	v[0] *= 10
}

func change2(v map[int]int) {
	v[10] = 99
}

func change3(v *int) {
	*v = 5
}
