package main

func main() {
	//cards := newDeck()
	//
	//hand, remainingCards := deal(cards, 5)
	//
	//hand.print()
	//fmt.Println()
	//remainingCards.print()

	cards := newDeck()
	cards.saveToFile("my_cards")
}
