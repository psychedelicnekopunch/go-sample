
package main


import (
	"fmt"
	"sort"
)


type Users struct {
	ID int
	Name string
}


type Lists [] Users


func (l Lists) Len() int {
	// Len is the number of elements in the collection.
	return len(l)
}


func (l Lists) Less(i, j int) bool {
	// Less reports whether the element with
	// index i should sort before the element with index j.
	return l[i].ID < l[j].ID
	// return i < j
}


func (l Lists) Swap(i, j int) {
	// Swap swaps the elements with indexes i and j.
	l[i], l[j] = l[j], l[i]
}


func main() {

	lists := Lists{}
	lists = append(lists, Users{ ID: 1, Name: "test1" })
	lists = append(lists, Users{ ID: 7, Name: "test7" })
	lists = append(lists, Users{ ID: 2, Name: "test2" })
	lists = append(lists, Users{ ID: 6, Name: "test6" })
	lists = append(lists, Users{ ID: 3, Name: "test3" })
	lists = append(lists, Users{ ID: 5, Name: "test5" })
	lists = append(lists, Users{ ID: 4, Name: "test4" })

	fmt.Println(lists)

	sort.Sort(lists)
	fmt.Println(lists)

	sort.Sort(sort.Reverse(lists))
	fmt.Println(lists)
}
