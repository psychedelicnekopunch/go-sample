
package main


import (
	"fmt"
)


type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Profile *string `json:"profile"`
	IsBan bool `json:"isBan"`
}


func main() {
	u := User{}
	fmt.Print(u, "\n")// {0 "" <nil> false}

	// Profile に代入する時
	profile := "プロフィール"
	u.Profile = &profile

	// Profile を使う時
	// nil チェックをしないとバグの原因になったりする
	if u.Profile != nil {
		fmt.Print(*u.Profile, "\n")// プロフィール
	}
}
