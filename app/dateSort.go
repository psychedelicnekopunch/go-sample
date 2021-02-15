package main

import (
	"fmt"
	"sort"
	"time"
)


func main() {
	// 2010.04.01〜2010.04.02, 2019.08.27, 2019.08.29〜2019.08.31, 2019.09.30
	dates := []string{ "2019.08.31", "2019.08.27", "2019.08.29", "2019.08.30", "2019.08.31", "2019.09.30", "2010.04.02", "2010.04.01" }

	type Date struct {
		unix int
		str string
	}

	// formatedMaps := map[int]string{}
	formatedMaps := map[int]Date{}

	// func Parse(layout, value string) (Time, error)
	// layout: Mon Jan 2 15:04:05 -0700 MST 2006
	// layout: Mon 01 02 15:04:05 -0700 MST 2006
	for _, date := range dates {
		t, err := time.Parse("2006.01.02", date)
		if err != nil {
			fmt.Print(err.Error())
			return
		}

		key := int(t.Unix())
		// fmt.Print(t.Format(time.UnixDate))
		// fmt.Print(t.Unix())
		// fmt.Print("\n")
		// formatedMaps[key] = date
		formatedMaps[key] = Date{ unix: key, str: date }
	}


	ints := []int{}
	for i, _ := range formatedMaps {
		ints = append(ints, i)
	}
	sort.Ints(ints)
	// fmt.Print(ints)
	// fmt.Print(formatedMaps)


	// formateds := []string{}
	formateds := []Date{}
	for _, i := range ints {
		formateds = append(formateds, formatedMaps[i])
	}
	// fmt.Print(formateds)


	res := ""
	lastIndex := len(ints) - 1
	skip := false

	for i, formated := range formateds {
		if i == lastIndex {
			res += formated.str
			continue
		}
		if i == 0 {
			res += formated.str
			if lastIndex == 0 {
				continue
			}
			if formateds[i + 1].unix - formated.unix <= 24*60*60 {
				res += "〜"
				skip = true
			} else {
				res += ", "
			}
			continue
		}
		if formateds[i + 1].unix - formated.unix <= 24*60*60 {
			if skip {
				continue
			}
			res += formated.str
			res += "〜"
			skip = true
			continue
		}
		res += formated.str
		res += ", "
		skip = false
	}

	fmt.Print(res)
}
