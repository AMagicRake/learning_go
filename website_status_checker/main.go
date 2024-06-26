package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkStatus(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url + " is not available.")
		c <- url
		return
	}
	fmt.Println(url + " is not available.")
	c <- url
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

	for l := range output {
		go func() {
			time.Sleep(2 * time.Second)
			checkStatus(l, output)
		}()
	}
}
