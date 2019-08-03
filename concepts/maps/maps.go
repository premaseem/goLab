package main

import "fmt"

type anim struct {
	color string
	age   int
}

func main() {

	var animals = make(map[string]int)

	animals["dog"] = 15

	fmt.Println(animals)
	fruits := map[string]string{
		"mango": "yellow",
		"aplle": "red",
	}

	iterMap(fruits)
	delete(fruits, "mango")
	fruits["apple"] = "green"
	iterMap(fruits)

}

func iterMap(m map[string]string) {

	for _, v := range m {
		fmt.Println(v)
	}
}
