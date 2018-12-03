package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	//alex := person{"Alex", "Anderson"}
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//var alex person
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"

	alex := person{
		firstName: "Alex",
		lastName: "Anderson",
		contact: contactInfo{
			email: "aa@email.com",
			zipCode: 123,
		},
	}

	fmt.Printf("%+v", alex)
}