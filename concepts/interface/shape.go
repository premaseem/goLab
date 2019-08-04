package main

import "fmt"

type shape interface {
	area() int
}

type square struct {
	side int
}

type rectangle struct {
	lenght int
	width  int
}

func measure(s shape) int {
	return s.area()
}

func (s square) area() int {
	return s.side * s.side
}

func (s rectangle) area() int {
	return s.lenght * s.width
}

func main() {
	squr := rectangle{lenght: 2, width: 5}
	rect := square{
		side: 10,
	}

	fmt.Println(shape.area(squr))
	fmt.Println(shape.area(rect))
	fmt.Println(measure(squr))
	fmt.Println(measure(rect))
}
