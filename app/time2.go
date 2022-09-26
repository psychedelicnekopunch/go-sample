package main

import (
	"fmt"
	"time"
)


func main() {

	d := "2022.11.28"
	loc, _ := time.LoadLocation("Asia/Tokyo")


	t, err := time.ParseInLocation("2006.01.02", d, loc)
	if err != nil {
		fmt.Print(err.Error(), "\n")
		return
	}
	s := t.Unix() + (60*60*24)
	fmt.Print(s, "\n")// 1669647600


	parse, err := time.Parse("2006.01.02", d)
	if err != nil {
		fmt.Print(err.Error(), "\n")
		return
	}
	t2 := parse.In(loc)
	s2 := t2.Unix() + (60*60*24)
	fmt.Print(s2, "\n")// 1669680000
}
