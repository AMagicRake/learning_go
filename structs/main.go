package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p person) toString() string {
	return p.firstName + " " + p.lastName
}

func main() {

	niel := person{"Niel", "Gow"}

	fmt.Println(niel.toString())
}
