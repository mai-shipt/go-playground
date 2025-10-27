package main

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
}

// fxn declaration explicitly returns a string
func newCard() string {
	return "Five of Diamonds"
}
