package models

import "fmt"

type Band struct {
	ID int
	Name string
}

func NewBands() Band {
	return Band{}
}

func (c Band) GetById(bandId int) interface{} {

	var bands []Band

	bands = append(bands, Band{1, "Red Hot Chili Peppers"})
	bands = append(bands, Band{2, "Prodigy"})
	bands = append(bands, Band{3, "Dodos"})

	fmt.Printf("bandId is %v\n", bandId)

	for _, band := range bands {
		if (bandId == band.ID) {
			return band
		}
	}

	return nil
}
