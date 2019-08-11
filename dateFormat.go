package main

import (
	"fmt"
	// "sort"
	"time"
)


func main() {

	// parsing time "2019/08/01 12:23:45": extra text:  12:23:45
	// dateStr := "2019/08/01 12:23:45"
	// t, err := time.Parse("2006/01/02", dateStr)

	// cannot parse "1" as "02"
	// dateStr := "2019/08/1"
	// t, err := time.Parse("2006/01/02", dateStr)
	// t, err := time.Parse("2006/01/2", dateStr)

	// month out of range
	// dateStr := "2019/8/01"
	// t, err := time.Parse("2006/01/02", dateStr)
	// t, err := time.Parse("2006/1/02", dateStr)

	// parsing time "2019/08/01" as "2006/Jan/02": cannot parse "08/01" as "Jan"
	// dateStr := "2019/08/01"
	// dateStr := "2019/Oct/01"
	// dateStr := "2019/OCT/01"
	// t, err := time.Parse("2006/Jan/02", dateStr)

	// layout: Mon Jan 2 15:04:05 -0700 MST 2006
	// layout: Mon 01 02 15:04:05 -0700 MST 2006

	// parsing time "2019.08.01 12時23分45秒" as "2006/01/02 15:04:05": cannot parse ".08.01 12時23分45秒" as "/"
	// dateStr := "2019.08.01 12時23分45秒"
	// t, err := time.Parse("2006/01/02 15:04:05", dateStr)
	// t, err := time.Parse("2006.01.02 15時04分05秒", dateStr)

	// string to date (Time)
	dateStr := "2019/08/01 12:23:45"
	t, err := time.Parse("2006/01/02 15:04:05", dateStr)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Print(t)
	fmt.Print("\n")

	// unix (int) to date (Time)
	unix := int64(1564662225)
	t = time.Unix(unix, 0)
	fmt.Print(t)
	fmt.Print("\n")

	// date (Time) to unix (int64)
	u := time.Now().Unix()
	fmt.Print(u)
	fmt.Print("\n")

	// date (Time) to string
	s := time.Now().Format("2006.01.02(Mon) 15:04:05")
	fmt.Print(s)
}
