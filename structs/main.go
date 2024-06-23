package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	contactInfo `json:"contactInfo"`
}

type contactInfo struct {
	Email        string `json:"emailAddress"`
	MobileNumber string `json:"mobileNumber"`
}

func (p person) toString() string {
	return p.FirstName + " " + p.LastName + " " + p.contactInfo.toString()
}

func (ci contactInfo) toString() string {
	return ci.Email + " " + ci.MobileNumber
}

func (p *person) updateName(newFirstName string) {
	p.FirstName = newFirstName
}

func main() {

	niel := person{"Niel", "Gow", contactInfo{"niel.gow@domain.com", "+4407123456789"}}
	meg := person{FirstName: "Megan", LastName: "Smith"}
	meg.contactInfo = contactInfo{Email: "meg.smith@domain.com", MobileNumber: "+467123456789"}

	meg.updateName("Meg")

	fmt.Println(niel.toString())
	fmt.Println(meg.toString())
	fmt.Printf("%+v\n", niel)

	// data := &person{"John", "Gow"}
	jsonNiel, err := json.MarshalIndent(niel, "", "    ")
	if err != nil {
		fmt.Println("could not marshall json")
	} else {
		fmt.Printf("json data: %s\n", jsonNiel)
	}

}
