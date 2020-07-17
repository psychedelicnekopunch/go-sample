
package main


import (
	"fmt"
	"math/rand"
	"time"
	"unicode/utf8"
)



type Users struct {
	ID int
	Name string
	Number int
}


type Random struct {
	MaxRange int
	MinRange int
	Value string
}


func NewRandom() *Random {
	r := new(Random)
	r.MaxRange = 8
	r.MinRange = 8
	r.Value = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return r
}


func (r *Random) Get() string {

	var source rand.Source
	source = rand.NewSource(time.Now().UnixNano())

	var str []rune
	var cnt int
	var res string

	str = []rune(r.Value)
	cnt = utf8.RuneCountInString(r.Value)

	for i := 0; i < r.MaxRange; i++ {
		if r.MinRange <= i {
			if rand.New(source).Intn(r.MaxRange - r.MinRange) == 0 {
				break
			}
		}
		key := rand.New(source).Intn(cnt)
		res = res + fmt.Sprintf("%c", str[key])
	}
	return res
}


func getUsers() (users []Users) {

	random := NewRandom()
	maxNumber := 4

	var source rand.Source = rand.NewSource(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		num := rand.New(source).Intn(maxNumber)
		user := Users{
			ID: (i + 1),
			Name: random.Get(),
			Number: num + 1,
		}
		users = append(users, user)
	}
	return users
}


func main() {

	users := getUsers()
	selectedUsersIndex := map[int]bool{}
	selectedUsers := []Users{}

	for _, user := range users {
		selectedUsersIndex[user.ID] = false
	}

	selected := 50
	cnt := len(selectedUsersIndex)

	if cnt < selected {
		selected = cnt
	}

	maxNumber := 4
	luckey := 0

	var source rand.Source = rand.NewSource(time.Now().UnixNano())

	for {
		if luckey == selected {
			break
		}
		index := rand.New(source).Intn(cnt)
		user := users[index]
		if selectedUsersIndex[user.ID] {
			continue
		}
		if maxNumber < user.Number {
			continue
		}
		if luckey + user.Number > selected {
			maxNumber--
			found := false
			for _, user := range users {
				if user.Number <= maxNumber {
					found = true
				}
			}
			if !found {
				break
			}
			continue
		}
		selectedUsers = append(selectedUsers, user)
		selectedUsersIndex[user.ID] = true
		luckey += user.Number
		fmt.Print(user.Number, "+")
	}

	fmt.Print("\n")

	for _, selectedUser := range selectedUsers {
		fmt.Print(selectedUser, "\n")
	}

	fmt.Print(len(selectedUsers), "\n")
	fmt.Println(luckey)
	// fmt.Println(selectedUsersIndex)
}
