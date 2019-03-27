
package main


import (
	"fmt"
	"regexp"
	"unicode/utf8"
)


type RegexpSample struct {}


func NewRegexpSample() *RegexpSample {
	return new(RegexpSample)
}


func (r *RegexpSample) Password() *regexp.Regexp {
	reg := `(?i)^[a-z0-9]+$`
	return regexp.MustCompile(reg)
}


func (r *RegexpSample) Email() *regexp.Regexp {
	reg := `(?i)^[0-9a-z].[0-9a-z-._]+@[0-9a-z].[0-9a-z-._]+[0-9a-z]+$`
	return regexp.MustCompile(reg)
}


func (r *RegexpSample) Image() *regexp.Regexp {
	reg := `(?i)(\.jpg|\.jpeg|\.gif|\.png)$`
	return regexp.MustCompile(reg)
}

func main() {
	reg := NewRegexpSample()

	regPassword := reg.Password()
	fmt.Print("> Password: ")
	fmt.Println(regPassword)
	fmt.Print(regPassword.MatchString(""), "\n")// false
	fmt.Print(regPassword.MatchString("r"), "\n")// true
	fmt.Print(regPassword.MatchString("r\n"), "\n")// false
	fmt.Print(regPassword.MatchString(" test"), "\n")// false
	fmt.Print(regPassword.MatchString("Success1"), "\n")// true
	fmt.Print(regPassword.MatchString("Failed-Word"), "\n")// false

	regEmail := reg.Email()
	fmt.Print("> Email: ")
	fmt.Println(regEmail)
	fmt.Print(regEmail.MatchString("psychedelic.nekopunch@gmail.com"), "\n")// true
	fmt.Print(regEmail.MatchString("psychedelic.nekopunch@yahoo.ne.jp"), "\n")// true
	fmt.Print(regEmail.MatchString("psychedelic.nekopunch@r.vodafone.ne.jp"), "\n")// true

	regImage := reg.Image()
	fmt.Print("> Image: ")
	fmt.Println(regImage)
	fmt.Print(regImage.MatchString("The Velvet Underground & Nico.jpg"), "\n")// true
	fmt.Print(regImage.MatchString("Nicojpg"), "\n")// false
	fmt.Print(regImage.MatchString("Nico.gifjpg"), "\n")// alse
}