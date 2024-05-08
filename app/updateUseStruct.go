
package main


import (
	"fmt"
)


type Person struct {
	Age int
}


func NewPerson() *Person {
	return new(Person)// ok
	// return &Person{}// ok
}


func (s *Person) Say() string {
	return "HELLO"
}


func main() {

	p := Person{ Age: 10 }// ok
	// p := &Person{ Age: 10 }// ok
	// p.Say()
	fmt.Print("p.Age " , p.Age, "\n")
	fmt.Print("p.Say() " , p.Say(), "\n")

	// say := Person{}.Say()// NG
	say := new(Person).Say()// ok
	// say := &Person{}.Say()// NG
	fmt.Print("Person{}.Say() " , say, "\n")

	age := Person{}.Age// ok
	// age := new(Person).Age// ok
	// age := &Person{}.Age// NG
	fmt.Print("Person{}.Age " , age, "\n")

	fmt.Print("NewPerson().Age " , NewPerson().Age, "\n")

	fmt.Print("-----\n")


	person := Person{ Age: 2 }

	change(person)
	fmt.Print("Person ", person, "\n")

	change2(&person)
	fmt.Print("Person ", person, "\n")



	m1 := make(map[int]Person)
	m1[0] = Person{}
	fmt.Print("m1[0].Age ", m1[0].Age, "\n")
	// fmt.Print("m1[0].Say() ", m1[0].Say(), "\n")// cannot call pointer method Say on Person
	for _, v := range m1 {
		fmt.Print("v.Say() ", v.Say(), "\n")
	}


	m2 := make(map[int]*Person)
	m2[0] = &Person{}
	fmt.Print("m2[0].Age ", m2[0].Age, "\n")
	fmt.Print("m2[0].Say() ", m2[0].Say(), "\n")
	for _, v := range m2 {
		fmt.Print("v.Say() ", v.Say(), "\n")
	}


	s := make([]Person, 1)
	s[0] = Person{}
	fmt.Print("s[0].Age ", s[0].Age, "\n")
	fmt.Print("s[0].Say() ", s[0].Say(), "\n")
	for _, v := range s {
		fmt.Print("v.Say() ", v.Say(), "\n")
	}


}


func change(v Person) {
	v.Age *= 10
	fmt.Print("change() v.Say() " , v.Say(), "\n")
}

func change2(v *Person) {
	v.Age *= 10
	fmt.Print("change2() v.Say() " , v.Say(), "\n")
}
