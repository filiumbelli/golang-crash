package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	// Updated version
	contactInfo
}
type contactInfo struct {
	email string
	zip   int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
func main() {
	alex := person{firstName: "Alex", lastName: "De Souza", contactInfo: contactInfo{zip: 1, email: "test@email.com"}}

	alexPointer := &alex
	alexPointer.changeEmailByReference("changed@email.com")
	alex.print()
	// or
	alex.changeEmailByReference("withoutAddress@email.com")
	alex.print()
	changeEmailByRef(&alex, "another@email.com")
	alex.print()
}

func changeNameByValue(p person) {
	p.lastName = "De Souza ByValue"
}
func changeEmailByRef(p *person, email string) {
	p.contactInfo.email = email
}
func (pointerToPerson *person) changeEmailByReference(email string) {
	(*pointerToPerson).contactInfo.email = email
}
