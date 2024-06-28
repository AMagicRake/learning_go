package main

import (
	"fmt"
	"net/http"
)

func checkStatus(url string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, " is not available.")
		return
	}
	fmt.Println(url, "is available.")
}

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		go checkStatus(link)
	}
}
