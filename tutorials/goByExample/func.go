package main

import "fmt"

func main() {
	sm, dff := sumAndProduct(5, 2, 5, 5)
	fmt.Println("sum is ", sm, " and product is ", dff)

}

func sumAndProduct(a ...int) (int, int) {
	fmt.Println(a)
	var sum int
	var prd int = 1
	for _, v := range a {
		sum += v
		prd *= v
	}
	return sum, prd
}
