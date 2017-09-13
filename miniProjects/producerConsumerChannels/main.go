package main

import "fmt"

func main() {
	fmt.Println("Using send only channels ")
	c := make(chan int, 10)
	c <- 420
	producer(c)
	consumer(c)



}

func producer(c chan <- int){
	c <- 1
	c <- 2
	c <- 3
	close(c)
}

func consumer(c <-chan int ){
	// reading from channel
	for v := range c {
		fmt.Printf("v = %v \n", v)
	}
}