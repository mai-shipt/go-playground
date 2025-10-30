package main

import (
	"fmt"
	"io"
	"os"
)

// Open a file and print out its content to standard output
func main() {
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}

// go build main.go to get executable so we can pass in an argument!
