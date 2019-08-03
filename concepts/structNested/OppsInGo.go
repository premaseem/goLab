package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p person) walk() {
	fmt.Println(p.firstName, "is walking")
}

func main() {
	var prem person
	prem.firstName = "prem"
	prem.lastName = "jain"
	prem.walk()

}
