package main

import     "fmt"
import "time"


type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo


var list = Todos{
Todo{Name:true, Name:"aseem"},
}

func kill()  {
	fmt.Print("hanuman")
}