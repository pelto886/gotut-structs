package main

import "fmt"

// regarding pointer:
// since we are using a pass by language
// and of how golang handles referenses
// value type variables need pointers to change inside a function
// and reference type variables does not.
// value types: int, float, string, bool, structs
// reference types: slices, maps, channels, pointers, functions

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "jim",
		lastName:  "party",
		contactInfo: contactInfo{
			email:   "jim.party@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Karl")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson *person) print() {
	fmt.Printf("%+v", pointerToPerson)
}
