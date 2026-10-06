package main

import (
	"fmt"
	"github.com/psychedelicnekopunch/go-sample/internal/domain"
)

func main() {
	// user := users.Get(2)
	users := domain.NewUsers()
	user := users.GetById(2)
	fmt.Printf("User is %v\n", user)

	bands := domain.NewBands()
	band := bands.GetById(1)
	fmt.Printf("Band is %v\n", band)
}
