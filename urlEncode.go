package main

import (
	"fmt"
	"net/url"
)


func main() {
	str := "❤️ 絵文字_ やゔぁぴぼ _ 空白 \" ' ? ! .xlsx .png ♡"
	fmt.Print(str, "\n\n")
	fmt.Print(url.PathEscape(str), "\n\n")
	fmt.Print(url.QueryEscape(str), "\n\n")
}
