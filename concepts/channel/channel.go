package main

import "fmt"

// Go routines

func main() {

	fmt.Println("Go Channels ")
	channel := make(chan string)

	go routine("non blocking 1", channel)

	go routine("non blocking", channel)

	message1 := <-channel
	fmt.Println(message1)

	message2 := <-channel
	fmt.Println(message2)

}

func routine(s string, c chan string) {
	for i := 0; i < 3; i++ {
		fmt.Println(i, s)
	}
	c <- s + "done"
}
