package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)


func main() {

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	f.SetCellValue("Sheet1", "A1", "日本語")
	f.SetCellValue("Sheet1", "A2", "にほんご")
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := f.SaveAs("./tmp/Book1.xlsx")
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Printf("created\n")
}
