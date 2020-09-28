package main

import (
	"fmt"
	"github.com/psychedelicnekopunch/go-sample/models"
)

func main() {
	// user := users.Get(2)
	users := models.NewUsers()
	user := users.GetById(2)
	fmt.Printf("User is %v\n", user)

	bands := models.NewBands()
	band := bands.GetById(1)
	fmt.Printf("Band is %v\n", band)
}
