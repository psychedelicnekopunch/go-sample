package main

import (
	"encoding/csv"
	"fmt"
	"os"
)


func main() {

	file, err := os.OpenFile("tmp/people.csv", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer file.Close()

	err = file.Truncate(0) // ファイルを空っぽにする(実行2回目以降用)
    if err != nil {
		fmt.Print(err.Error())
		return
	}

	writer := csv.NewWriter(file)
	writer.Write([]string{"Alice", "20"})
	writer.Write([]string{"日本", "21"})
	writer.Write([]string{"カナ", "22"})
	writer.Flush()

	fmt.Printf("created\n")
}
