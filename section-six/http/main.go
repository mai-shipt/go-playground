package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Go and make an HTTP GET request and handle response and error
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Equivalent to bottom block of code
	io.Copy(os.Stdout, resp.Body)

	//// Create an empty byte slice with spaces for 99,999 elements
	//bs := make([]byte, 99999)
	//// Now we can read the actual doctype HTML response
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))
}
