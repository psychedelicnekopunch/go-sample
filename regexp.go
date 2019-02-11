
package main


import (
	"fmt"
	"regexp"
)


type Regexp struct {}


func NewRegexp() *Regexp {
	return new(Regexp)
}


func (r *Regexp) Password() *regexp.Regexp {
	return regexp.MustCompile(`(?i)^[a-z0-9]+$`)
}


func (r *Regexp) Email() *regexp.Regexp {
	return regexp.MustCompile(`(?i)^[0-9a-z].[0-9a-z-._]+@[0-9a-z].[0-9a-z-._]+[0-9a-z]+$`)
}


func main() {
	reg := NewRegexp()

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
}