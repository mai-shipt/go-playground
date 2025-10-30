package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// Go and make an HTTP GET request and handle response and error
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// Equivalent to bottom block of code
	io.Copy(lw, resp.Body)

	//// Create an empty byte slice with spaces for 99,999 elements
	//bs := make([]byte, 99999)
	//// Now we can read the actual doctype HTML response
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))
}

// Custom implementation of the Writer interface: https://pkg.go.dev/io#Writer
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}
