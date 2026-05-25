package main


import (
	"encoding/json"
	"fmt"
	"strings"
)


func main() {

	type Test struct {
		ID int `json:"id"`
		Title string `json:"title"`
		OnSale bool `json:"onSale"`
	}


	var Test2 struct {
		ID int `json:"id"`
		Title string `json:"title"`
		OnSale bool `json:"onSale"`
	}

	test := Test{
		ID: 1,
		Title: "test",
		OnSale: true,
	}

	// Encode
	b, err := json.Marshal(test)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}
	fmt.Printf("%s\n", string(b))//{"id":1,"title":"test","onSale":true}

	// Decode
	docTest := Test{}
	doc := json.NewDecoder(strings.NewReader(string(b)))
	doc.Decode(&docTest)
	fmt.Print(docTest, "\n")

	// Decode2
	doc2 := json.NewDecoder(strings.NewReader(string(b)))
	doc2.Decode(&Test2)
	fmt.Print(Test2, "\n")
}
