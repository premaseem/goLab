package main

import "fmt"

type Person struct {
	name string
	age int
	snn string
	Address
}

type Address struct{
	street1 string
	street2 string
	city    string
	state   string
	Zip
}

// Zip defines a Zip code (eg: 12345-1105)
type Zip struct {
	code  string
	pobox string
}

func main(){
	var people map[string] Person  // nil map

	people = make(map[string]Person)

	var p Person  // declaration
	p = Person{"Mary", 35, "011-13-4567",
		Address{"1 First St", "we", "Brooklyn", "NY",
			Zip{"11212", "333"}},
	}

	people[p.name] = p

	for _, p = range people{
		fmt.Printf("%v lives in %v, %v \n", p.name, p.city, p.code  )
	}

	//output seems nested or contains
	// Mary lives in {1 First St we Brooklyn NY {11212 333}}, {11212 333}

	//output seems linear
	//Mary lives in Brooklyn, 11212


}


