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

	// &variable give the memory address the variable is pointing at
	// p := &joe
	//p.updateFirstName("Mike")
	joe.updateFirstName("Mike")
	joe.print()
}

// No pointer's needed for a slice, differ for a struct
// There's two kind of data in go, value types and reference types.
// References type is already a pointer by itself, so you don't need to use pointer for :
// slices, maps, channels, pointers, functions.
// But for all values types, you have to, eg :
// int, float, string, bool, structs ...

func (p *person) updateFirstName(firstname string) {
	// *p give the value of the memory address it points
	// (*p).firstName = firstname
	p.firstName = firstname
}

func (p person) print() {
	fmt.Println(p)
}
