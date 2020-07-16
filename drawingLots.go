
package main


import (
	"fmt"
	"math/rand"// https://golang.org/pkg/math/rand/
	"time"
	"unicode/utf8"
)



type Users struct {
	ID int
	Name string
}


type Random struct {
	MaxRange int
	MinRange int
	Value string
}


func NewRandom() *Random {
	r := new(Random)
	// r.MaxRange = 50
	// r.MinRange = 40
	r.MaxRange = 8
	r.MinRange = 8
	// r.Value = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r.Value = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"
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


func main() {

	random := NewRandom()
	users := []Users{}
	selectedUsersIndex := map[int]bool{}
	selectedUsers := []Users{}


	for i := 0; i < 100; i++ {
		user := Users{
			// ID: (i + 1) * 17,
			ID: (i + 1),
			Name: random.Get(),
		}
		users = append(users, user)
		selectedUsersIndex[i] = false
	}

	selected := 50
	cnt := len(selectedUsersIndex)

	if len(selectedUsersIndex) < selected {
		selected = len(selectedUsersIndex)
	}

	var source rand.Source = rand.NewSource(time.Now().UnixNano())

	for {
		if len(selectedUsers) == selected {
			break
		}
		index := rand.New(source).Intn(cnt)
		if selectedUsersIndex[index] {
			continue
		}
		selectedUsers = append(selectedUsers, users[index])
		selectedUsersIndex[index] = true
	}

	/*for userID, i := range selectedUsersIndex {
		fmt.Println(userID, i)
	}*/
	for _, selectedUser := range selectedUsers {
		fmt.Print(selectedUser, "\n")
	}

	fmt.Print(len(selectedUsers), "\n")

	fmt.Println(users)
	// fmt.Println(selectedUsersIndex)
}
