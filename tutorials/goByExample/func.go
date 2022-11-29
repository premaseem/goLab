package main

import "fmt"

func main() {
	sm, dff := sumAndMultiply(5, 2, 5, 5)
	fmt.Println("sum is ", sm, " and diff is ", dff)

}

func sumAndMultiply(a ...int) (int, int) {
	fmt.Println(a)
	var sum int
	var dff int
	for _, v := range a {
		sum += v
		dff -= v
	}
	return sum, dff
}
