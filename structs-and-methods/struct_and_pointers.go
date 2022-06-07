package main

import (
	"fmt"
	"strings"
)

type Person struct { // struct definition
	firstName string
	lastName  string
}

func upPerson(p *Person) { // function using struct as a parameter
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// 1- struct as a value type:
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	// 2 - struct as a pointer:
	pers2 := new(Person)
	pers2.firstName = "Robert"
	pers2.lastName = "Langdon"
	(*pers2).lastName = "Langdon" // this is also valid
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	// 3 - struct as a literal:
	pers3 := &Person{"Richard", "Clintown"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}
