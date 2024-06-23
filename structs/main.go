package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Comms     contactInfo `json:"contactInfo"`
}

type contactInfo struct {
	Email        string `json:"emailAddress"`
	MobileNumber string `json:"mobileNumber"`
}

func (p person) toString() string {
	return p.FirstName + " " + p.LastName + " " + p.Comms.toString()
}

func (ci contactInfo) toString() string {
	return ci.Email + " " + ci.MobileNumber
}

func main() {

	niel := person{"Niel", "Gow", contactInfo{"niel.gow@domain.com", "+4407123456789"}}
	meg := person{FirstName: "Megan", LastName: "Smith"}
	meg.Comms = contactInfo{Email: "meg.smith@domain.com", MobileNumber: "+467123456789"}

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
