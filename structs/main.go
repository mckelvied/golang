package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

func main() {
	alex := person{"Alex", "Anderson", 25, contactInfo{"alex@example.com", 12345}}
	fmt.Println(alex)
	p := person{
		firstName: "Alice",
		lastName:  "Smith",
		age:       30,
		contact:   contactInfo{"alice@example.com", 67890},
	}
	fmt.Println(p)

	var bob person
	bob.firstName = "Bob"
	bob.lastName = "Johnson"
	bob.age = 40
	bob.contact.email = "bob@example.com"
	bob.contact.zipCode = 54321
	fmt.Println(bob)
	bob.updateName("Robert")
	bob.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
