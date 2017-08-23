package main

import "fmt"

func Sum(a int, b int) (total int) {
	return a + b
}

func main() {
	x := Sum(2, 3)
	fmt.Print("is the toal sum ", x)
}
