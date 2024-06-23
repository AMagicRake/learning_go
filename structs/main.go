package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (p person) toString() string {
	return p.FirstName + " " + p.LastName
}

func main() {

	niel := person{"Niel", "Gow"}
	meg := person{FirstName: "Megan", LastName: "Smith"}

	fmt.Println(niel.toString())
	fmt.Println(meg.toString())
	fmt.Printf("%+v\n", niel)

	// data := &person{"John", "Gow"}
	jsonNiel, err := json.Marshal(niel)
	if err != nil {
		fmt.Println("could not marshall json")
	} else {
		fmt.Printf("json data: %s\n", jsonNiel)
	}

}
