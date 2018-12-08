package main

import "fmt"

type contactInfo struct {
	email string
	phone string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	joe := person{
		firstName: "Joe",
		lastName:  "Mann",
		contactInfo: contactInfo{
			email: "joe@mann.fr",
			phone: "0909090909",
		},
	}

	p := &joe
	p.updateFirstName("Mike")
	joe.print()
}

func (p *person) updateFirstName(firstname string) {
	(*p).firstName = firstname
}

func (p person) print() {
	fmt.Println(p)
}
