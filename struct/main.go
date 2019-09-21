package main

import "fmt"

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
		lastName:  "parti",
		contactInfo: contactInfo{
			email:   "jim@gmail.xmo",
			zipCode: 88888,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("jkimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
