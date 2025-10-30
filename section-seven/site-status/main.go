package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		// This will execute, but only first URL is checked... need a mechanism to block and wait!
		go checkLink(link, c)
	}

	fmt.Println(<-c)
}

// checkLink checks each URL and see if it is responding to traffic or not
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down..."
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep it's up!"
}
