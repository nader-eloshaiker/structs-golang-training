package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {

	alex := person{
		firstName: "Alex",
		lastName: "Anderson",
		contactInfo: contactInfo{
			email: "aa@email.com",
			zipCode: 123,
		},
	}

	alexPointer := &alex
	alexPointer.updateName("Alexar")
	alex.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)

}