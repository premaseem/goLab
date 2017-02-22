package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func changeName(p Person) {
	p.firstName = "Bad"
	fmt.Print("inside method with local scope on copy")
	fmt.Println(p)
}

func changeNameWithPointer(p *Person) {
	p.firstName = "Bad"
	fmt.Print("inside method with local scope on pointer")
	fmt.Println(p)
}

func main() {
	person := Person{
		firstName: "Good",
		lastName:  "person",
	}
	fmt.Println("Original value of person", person)
	changeName(person)
	fmt.Print("Outside method with original variable passed as copy")
	fmt.Println(person)

	fmt.Println("\n\n========Pass by pointer will modify values =======")

	fmt.Println("Outside method with original variable", person)
	changeNameWithPointer(&person)
	fmt.Println("Outside method with original variable passed as reference (which is now changed )")
	fmt.Println(person)
}
