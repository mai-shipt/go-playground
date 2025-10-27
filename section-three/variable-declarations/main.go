package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"	// declare and init
	card := "Ace of Spades" // shorthand declare and init (type inferred)
	card = "Ace of Clubs"   // re-assign

	fmt.Println(card)
}
