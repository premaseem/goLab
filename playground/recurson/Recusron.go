package main

func main() {
	println(factorial(9))
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}else {
		return n * factorial(n - 1)
	}
}