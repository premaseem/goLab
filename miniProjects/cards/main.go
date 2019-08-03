package main

import "fmt"

type address struct {
	line1 string
	line2 string
}

type person struct {
	firstName string
	lastName  string
	contact   address
}

func (pp *person) updateName(name string) {
	(*pp).firstName = name
}

func main() {
	// positional argument
	// prem := person{"aseem", "Jain"}
	// named argument or dictionary type

	// preston := address{
	// 	line1: "4414",
	// 	line2: "sat",
	// }

	// sony := person{
	// 	firstName: "prem",
	// 	lastName:  "jain",
	// 	contact:   preston,
	// }

	var meera person
	meera.firstName = "swasti"
	meera.lastName = " jain"
	meera.contact.line1 = "4114 medical dr"
	meera.contact.line2 = "san antonio"
	meeraAdd := &meera
	meeraAdd.updateName("mani")

	fmt.Printf("%+v", meera)
	// fmt.Printf("%+v", *meeraAdd)
	fmt.Print(*meeraAdd)
	fmt.Print(*&meera)

	// fmt.Print(prem)
	// fmt.Print(sony)
}
