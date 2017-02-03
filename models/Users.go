package models

import "fmt"

type User struct {
	ID int
	Name string
}

func NewUsers() User {
	return User{}
}

func (c User) GetById(userId int) interface{} {

	var users []User

	users = append(users, User{1, "Anthony"})
	users = append(users, User{2, "Tom"})
	users = append(users, User{3, "Masa"})

	fmt.Printf("userId is %v\n", userId)

	for _, user := range users {
		if (userId == user.ID) {
			return user
		}
	}

	return nil
}