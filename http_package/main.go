package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

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

	io.Copy(os.Stdout, resp.Body)
}
