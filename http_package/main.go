package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("https://www.google.co.uk")
	if err != nil {
		log.Fatal(err)
	}
	// bs := make([]byte, 99999)
	// bytes, err := resp.Body.Read(bs)
	// if bytes > 0 && err != nil {
	// 	fmt.Println(string(bs))
	// }
	// respBody, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(respBody))

	io.Copy(logWriter{}, resp.Body)
}
