package main

import (
	"fmt"
	"time"
)


func main() {

	var unix int64 = 1656342000 + (60 * 60 * 18) + (60 * 59) + 59

	loc, _ := time.LoadLocation("Asia/Tokyo")


	t := time.Unix(unix, 0).In(loc)
	s := t.Format("2006.01.02 15:04:05")
	fmt.Print(s, "\n")


	s2 := t.Format("2006.01.02 15:04")
	// fmt.Print(s2, "\n")


	t2, err := time.ParseInLocation("2006.01.02 15:04", s2, loc)
	if err != nil {
		fmt.Print(err.Error(), "\n")
		return
	}
	fmt.Print(unix, "\n")
	fmt.Print(t2.Unix(), "\n")


	s3 := t2.Format("2006.01.02 15:04:05")
	fmt.Print(s3, "\n")
}
