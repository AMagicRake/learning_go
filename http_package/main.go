package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.google.co.uk/search?q=reddit")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
