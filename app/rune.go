
package main


import (
	"fmt"
	// "strconv"
	"unicode/utf8"
)


func limit(str string) string {
	return string([]rune(str)[:30])
	// return string(utf8.RuneCountInString(str)[:15])
}


func main() {

	str1 := "abc, 日本語です。"
	fmt.Printf("%s (%d Words)(%d Byte)\n", str1, utf8.RuneCountInString(str1), len(str1))
	str1 = limit(str1)
	fmt.Printf("%s (%d Words)(%d Byte)\n", str1, utf8.RuneCountInString(str1), len(str1))


	str2 := "abc, 日本語です。 文字数制限LIMIT abc, 日本語です。 文字数制限LIMIT abc, 日本語です。 文字数制限LIMIT"
	fmt.Printf("%s (%d Words)(%d Byte)\n", str2, utf8.RuneCountInString(str2), len(str2))
	str2 = limit(str2)
	fmt.Printf("%s (%d Words)(%d Byte)\n", str2, utf8.RuneCountInString(str2), len(str2))


	str3 := "abc, 日本語です 12345。 ✋　😄 💖 文字数制限LIMIT abc, 日本語です。 文字数制限LIMIT abc, 日本語です。 文字数制限LIMIT"

	arr1 := []rune(str1)
	fmt.Println(arr1)

	arr3 := []rune(str3)
	fmt.Println(arr3)

	for _, v := range arr1 {
		fmt.Print(string(v), "|")
	}

	for _, v := range arr3 {
		fmt.Print(string(v), "|")
	}

	limitStr3 := ""
	for i, v := range arr3 {
		if i == 29 {
			break
		}
		limitStr3 = limitStr3 + string(v)
	}
	fmt.Print("\n==========\n", limitStr3, "\n==========\n")
}
