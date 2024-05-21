package main

import (
	"fmt"
	// "sort"
	"time"
)


func main() {
	// UTC: yyyy-MM-dd HH:mm:ss.SSS
	t := time.Now().UTC()
	// fmt.Print("\n", t, "\n")
	fmt.Print("\n", t.Unix(), "\n")
	fmt.Print("\n", t.Format("2006.01.02(Mon) 15:04:05.000"), "\n")
	fmt.Print("\n", t.Nanosecond(), "\n")
	fmt.Print("\n", time.Now().UTC().Format("2006.01.02(Mon) 15:04:05.000"), "\n")
}
