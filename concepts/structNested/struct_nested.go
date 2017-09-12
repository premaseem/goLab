package main


type Person struct {
	name string
	age int
	snn string
	address Address
}

type Address struct{
	street1 string
	street2 string
	city    string
	state   string
	zip Zip
}

// Zip defines a Zip code (eg: 12345-1105)
type Zip struct {
	code  string
	pobox string
}



