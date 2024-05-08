
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
	// s := make([]int, 2)// panic: runtime error: index out of range [3] with length 3
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
	change1(s)
	fmt.Print("change1(s) " , s, "\n")

	change2(s)
	fmt.Print("change2(s) " , s, "\n")

	change3(&s)
	fmt.Print("change3(&s) " , s, "\n")

	change4(s)
	fmt.Print("change4(s) " , s, "\n")

	change5(s)
	fmt.Print("change5(s) " , s, "\n")

	change6(&s)
	fmt.Print("change6(&s) " , s, "\n")

	for i, v := range s {
		// v = 5
		// s[i] = 5
		// change7(&v)
		change7(&s[i])
		fmt.Print("for: i " , i, ", v " , v, "\n")
	}
	fmt.Print("for " , s, "\n")

	change8(s)
	fmt.Print("change8(s) " , s, "\n")


	change9(&s)
	fmt.Print("change9(s) " , s, "\n")

	change4(s)
	fmt.Print("change4(s) " , s, "\n")
}


func change1(v []int) {
	v[0] = 10
}

func change2(v []int) {
	v = append(v, 99)
}

func change3(v *[]int) {
	*v = append(*v, 100)
}

func change4(v []int) {
	v[1] = 20
	v = append(v, 999)
}

func change5(v []int) {
	v = append(v, 9999)
	v[2] = 30
}

func change6(v *[]int) {
	// *v = append(*v, 200)
	copy := *v
	*v = append(*v, 200)
	// copy = append(copy, 200)// NG
	// *v[3] = 300// invalid operation: cannot index v (variable of type *[]int)
	copy[3] = 300
}

func change7(v *int) {
	*v = 5
}


func change8(v []int) {
	// v[0] = 100
	// for i := 0; i < 5; i++ {
	// 	v = append(v, i)
	// }
	v = append(v, 0)
	v = append(v, 0)
	v[0] = 100
}


func change9(v *[]int) {
	copy := *v
	// copy[0] = 1000
	// for i := 0; i < 5; i++ {
	// 	*v = append(*v, i)
	// }
	*v = append(*v, 0)
	*v = append(*v, 0)
	copy[0] = 1000
}
