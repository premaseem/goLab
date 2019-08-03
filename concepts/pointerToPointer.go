package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name
	namePointer1 := &namePointer
	namePointer2 := &namePointer1

	fmt.Println(&namePointer)
	fmt.Println(&namePointer1)
	fmt.Println(&namePointer2)
	printPointer(namePointer)

}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
