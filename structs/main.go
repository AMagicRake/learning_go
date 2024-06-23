package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	niel := person{"Niel", "Gow"}

	fmt.Println(niel.firstName, niel.lastName)
}
