
package main


import (
	"fmt"
	"strconv"
	"unicode/utf8"
)


type Users struct {
	ID int
	Name string
}



func (u *Users) CountName() int {
	return utf8.RuneCountInString(u.Name)
}


func main() {
	str := "abc, 日本語です。"
	fmt.Printf("%s (%d Byte)\n", str, len(str))// 23
	fmt.Printf("%s (%d Words)\n", str, len([]rune(str)))// 11
	fmt.Printf("%s (%d Words)\n", str, utf8.RuneCountInString(str))// 11

	var slice []Users
	slice = []Users{
		{ID: 1, Name: "test"},
		{ID: 2, Name: "test2"},
	}

	fmt.Print(slice, " (" + strconv.Itoa(len(slice)) + " values)\n")
	fmt.Printf("%s (%d Words)\n", slice[0].Name, slice[0].CountName())

	for i, value := range slice {
		fmt.Printf("%d: %s (%d Words)\n", i, value.Name, value.CountName())
	}
}