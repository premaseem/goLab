package main

import "fmt"

// Go routines

func main() {

	fmt.Println("Go Routines ")
	go routine("non blocking 1")

	routine("sequential")

	go routine("non blocking 2")

	fmt.Scanln()
}

func routine(s string) {
	for i := 0; i < 50; i++ {
		fmt.Println(i, s)
	}

}
