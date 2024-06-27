package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// fmt.Println(os.Args)
	fileName := os.Args[1]
	fmt.Println(fileName)

	myFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, myFile)
}
