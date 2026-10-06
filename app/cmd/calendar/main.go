
package main


import (
	"errors"
	"flag"
	"fmt"
	"time"
	"strconv"
	"sort"
)


type Calendar struct {
	Location *time.Location
	Selected struct {
		Time time.Time
	}
	Current struct {
		Time time.Time
	}
	Display struct {
		Time time.Time
		Lists CalendarLists
	}
	StartWeekday time.Weekday
}


type CalendarList struct {
	Time time.Time
	IsSelectedMonth bool
	IsToday bool
	Selected bool
}


type CalendarLists []CalendarList


func (c CalendarLists) Len() int {
	// Len is the number of elements in the collection.
	return len(c)
}


func (c CalendarLists) Less(i, j int) bool {
	// Less reports whether the element with
	// index i should sort before the element with index j.
	return (c[i].Time.Unix() < c[j].Time.Unix())
}


func (c CalendarLists) Swap(i, j int) {
	// Swap swaps the elements with indexes i and j.
	c[i], c[j] = c[j], c[i]
}


func NewCalendar(year int, month int, day int) *Calendar {

	loc, _ := time.LoadLocation("Asia/Tokyo")

	c := new(Calendar)

	c.Location = loc
	// c.StartWeekday = time.Sunday
	c.StartWeekday = time.Monday
	c.Selected.Time = time.Unix(time.Now().Unix(), 0)
	c.Current.Time = c.Selected.Time
	c.Display.Time = c.Selected.Time

	if year != 0 && month != 0 && day != 0 {
		selectedDate := fmt.Sprintf("%d/%s/%s", year, c.fillZero(month), c.fillZero(day))
		selectedTime, err := time.ParseInLocation("2006/01/02", selectedDate, c.Location)
		if err != nil {
			fmt.Print(err.Error())
			return c
		}
		c.Selected.Time = selectedTime
		c.Display.Time = selectedTime
	}

	if err := c.update(); err != nil {
		fmt.Print(err.Error())
	}

	return c
}


func (c *Calendar) monthToIntWithErr(month time.Month) (i int, err error) {
	switch month {
	case time.January:
		return 1, nil
	case time.February:
		return 2, nil
	case time.March:
		return 3, nil
	case time.April:
		return 4, nil
	case time.May:
		return 5, nil
	case time.June:
		return 6, nil
	case time.July:
		return 7, nil
	case time.August:
		return 8, nil
	case time.September:
		return 9, nil
	case time.October:
		return 10, nil
	case time.November:
		return 11, nil
	case time.December:
		return 12, nil
	}
	return 0, errors.New("unexpected error")
}


func (c *Calendar) monthToInt(month time.Month) int {
	i, _ := c.monthToIntWithErr(month)
	return i
}


func (c *Calendar) fillZero(i int) string {
	if i < 10 {
		return fmt.Sprintf("0%d", i)
	}
	return fmt.Sprintf("%d", i)
}


func (c *Calendar) update() (err error) {

	firstDate := fmt.Sprintf("%d/%s/01", c.Display.Time.Year(), c.fillZero(c.monthToInt(c.Display.Time.Month())))
	firstTime, err := time.ParseInLocation("2006/01/02", firstDate, c.Location)
	if err != nil {
		return err
	}

	c.Display.Lists = CalendarLists{}
	preLists := CalendarLists{}
	lists := CalendarLists{}

	for i := 0; i < 31; i++ {
		t := time.Unix(firstTime.Unix() + int64(60*60*24*i), 0)
		if c.Display.Time.Year() == t.Year() && c.Display.Time.Month() == t.Month() {
			lists = append(lists, CalendarList{
				Time: t,
				IsSelectedMonth: (c.Selected.Time.Year() == t.Year() && c.Selected.Time.Month() == t.Month()),
				IsToday: (c.Current.Time.Year() == t.Year() && c.Current.Time.Month() == t.Month() && c.Current.Time.Day() == t.Day()),
				Selected: (c.Selected.Time.Day() == t.Day()),
			})
		}
	}

	if lists[0].Time.Weekday() != c.StartWeekday {
		for i := -1; i >= -7; i-- {
			t := time.Unix(firstTime.Unix() + int64(60*60*24*i), 0)
			preLists = append(preLists, CalendarList{
				Time: t,
				IsSelectedMonth: (c.Selected.Time.Year() == t.Year() && c.Selected.Time.Month() == t.Month()),
				IsToday: (c.Current.Time.Year() == t.Year() && c.Current.Time.Month() == t.Month() && c.Current.Time.Day() == t.Day()),
				Selected: (c.Selected.Time.Day() == t.Day()),
			})
			if t.Weekday() == c.StartWeekday {
				break
			}
		}
		// sort.Sort(sort.Reverse(preLists))
		sort.Sort(preLists)
		// fmt.Print(preLists)
		for _, preList := range preLists {
			c.Display.Lists = append(c.Display.Lists, preList)
		}
	}

	for _, list := range lists {
		c.Display.Lists = append(c.Display.Lists, list)
	}

	if len(c.Display.Lists) < (7*5) {
		startIndex := len(c.Display.Lists)
		for i := startIndex; i < (7*5); i++ {
			t := time.Unix(c.Display.Lists[0].Time.Unix() + int64(60*60*24*i), 0)
			c.Display.Lists = append(c.Display.Lists, CalendarList{
				Time: t,
				IsSelectedMonth: (c.Selected.Time.Year() == t.Year() && c.Selected.Time.Month() == t.Month()),
				IsToday: (c.Current.Time.Year() == t.Year() && c.Current.Time.Month() == t.Month() && c.Current.Time.Day() == t.Day()),
				Selected: (c.Selected.Time.Day() == t.Day()),
			})
			// fmt.Printf("%d\n", i)
		}
	}

	return nil
}




func main() {
	flag.Parse()
	args := flag.Args()

	var y, m, d int

	if len(args) == 3 {
		year, _ := strconv.Atoi(args[0])
		y = year

		month, _ := strconv.Atoi(args[1])
		m = month

		day, _ := strconv.Atoi(args[2])
		d = day
	}

	c := NewCalendar(y, m, d)
	for _, list := range c.Display.Lists {
		fmt.Print(list)
		fmt.Print("\n")
	}
}
