
package main


import (
	"fmt"
	"time"
)

func main() {
	var u int64 = 1550200493// - (60*60*24)
	fmt.Println(getList(u))
}


// type Unixs struct {
// 	Start int64
// 	Finish int64
// }


// func getList(baseUnix int64) (lists []Unixs) {
func getList(baseUnix int64) (lists []struct{
	Start int64
	Finish int64
}) {

	lists = []struct {
		Start int64
		Finish int64
	}{}

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return lists
	}

	var baseTime time.Time

	for {
		t := time.Unix(baseUnix, 0).In(loc)
		// fmt.Print(t.Day(), "\n")
		if t.Day() == 1 || t.Day() == 15 {
			// fmt.Print(t.Day())
			fmt.Print(t.Year(), t.Month(), t.Day(), "\n")
			str := fmt.Sprintf("%v/%v/%v 00:00:00", t.Year(), t.Month(), t.Day())
			// baseTime, err = time.ParseInLocation("2006/01/02 15:04:05", str, loc)
			baseTime, err = time.ParseInLocation("2006/January/2 15:04:05", str, loc)
			if err != nil {
				fmt.Print(err.Error(), "\n")
				return lists
			}
			break
		}
		baseUnix = baseUnix - (60*60*24)
	}

	// fmt.Print(baseTime.Unix(), "\n")

	currentUnix := time.Now().Unix()

	startTime := baseTime
	finishTime := time.Unix(baseTime.Unix() + (60*60*24) , 0).In(loc)

	for {
		if finishTime.Day() == 1 || finishTime.Day() == 15 {
			if currentUnix < finishTime.Unix() {
				break
			}
			lists = append(lists, struct{
				Start int64
				Finish int64
			}{
				Start: startTime.Unix(),
				Finish: finishTime.Unix() - 1,
			})
			startTime = finishTime
			finishTime = time.Unix(finishTime.Unix() + (60*60*24) , 0).In(loc)
			continue
		}
		finishTime = time.Unix(finishTime.Unix() + (60*60*24) , 0).In(loc)
	}

	return lists
}
