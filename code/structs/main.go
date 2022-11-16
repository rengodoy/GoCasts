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
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// jimPointer := &jim
	// fmt.Printf("%p", &jim)
	// jimPointer.updateName("renato")
	jim.updateName("jimmy")
	jim.print()

	name := []string{"Bill"}
	updateValue(name)
	fmt.Println(name)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func updateValue(n []string) {
	n[0] = "Alex"
}
