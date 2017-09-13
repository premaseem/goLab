package main

import "fmt"

func main() {

	var c chan int // declaration
	fmt.Printf("c = %v", c)

	// check length, it is unbuffered so len will be 0
	fmt.Printf("chan len = %v", len(c))


	c = make(chan int, 2)
	// check length, it is unbuffered so len will be 0
	fmt.Printf("chan len = %v", len(c))

	// writing to channel
	c <- 1
	c <-786
	//c <- 6162

	fmt.Printf("chan len = %v", len(c))

	var v = <- c
	fmt.Println("len of c: ", len(c))
	fmt.Printf("v = %v\n", v)

}


