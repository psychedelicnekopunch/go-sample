
package main


import (
	"fmt"
)


func main() {
	// ----- //
	// s := []int{0, 0, 0}
	// ----- //
	// s := []int{}
	// s = append(s, 1)
	// s = append(s, 2)
	// s = append(s, 3)
	// ----- //
	// s := make([]int, 1)
	s := make([]int, 3)
	// s := make([]int, 3, 3)
	// s = append(s, 1000)
	// s[0] = 1
	// s[1] = 2
	// s[2] = 3
	// for i, _ := range s {
	// 	s[i] = i + 1
	// }
	// ----- //
	// s := make([]int, 0)
	// s = append(s, 1)
	// s = append(s, 2)
	// s = append(s, 3)
	// ----- //
	change(s)
	fmt.Print("change(s) " , s, "\n")

	change2(s)
	fmt.Print("change2(s) " , s, "\n")

	change3(s)
	fmt.Print("change3(s) " , s, "\n")

	change4(s)
	fmt.Print("change4(s) " , s, "\n")

	for i, v := range s {
		// v = 5
		// s[i] = 5
		// change5(&v)
		change5(&s[i])
		fmt.Print("for: i " , i, ", v " , v, "\n")
	}
	fmt.Print("for " , s, "\n")
}


func change(v []int) {
	v[0] = 10
}

func change2(v []int) {
	v = append(v, 99)
}

func change3(v []int) {
	v[1] = 20
	v = append(v, 999)
}

func change4(v []int) {
	v = append(v, 9999)
	v[2] = 30
}

func change5(v *int) {
	*v = 5
}
