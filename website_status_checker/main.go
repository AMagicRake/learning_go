package main

import (
	"fmt"
	"net/http"
)

func checkStatus(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		c <- url + " is not available."
		return
	}
	c <- url + " is available."
}

func main() {

	output := make(chan string)
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		go checkStatus(link, output)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-output)
	}
}
