package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Type of 'deck' has all features of a slice of strings
type deck []string

// newDeck() create and return a list of playing cards (aka a deck: an array of strings)
// Acts like a constructor, returns a value instance of 'deck'
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// Use blank identifier '_' to ignore index
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// print() log out the contents of a deck of cards
// This method is related to "type" deck
// Any variable of type "deck" have access to this "print" method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal() create a 'hand' of cards
// Go supports multiple return types with '(deck, deck)' syntax
func deal(d deck, handSize int) (deck, deck) {
	// first: return everything up to, not including, handSize
	// second: return everything after handSize, including
	return d[:handSize], d[handSize:]
}

// toString() take a deck and turn it into a string
// Take a deck, convert to []string, then take each string, convert to []byte
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// saveToFile() saves deck to local hard drive
func (d deck) saveToFile(filename string) error {
	// 0666 permission to let anyone read/write this file
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// newDeckFromFile() reads a deck from local hard drive
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option 1: log the error and return a call to newDeck()
		// Option 2: log the error and quit the program entirely
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// shuffle() shuffles all the cards and randomize the order in a deck
func (d deck) shuffle() {
	// Use 'time.Now().UnixNano()' to generate a new int64 number everytime program starts
	// Then use that as a seed to generate a new Source object
	source := rand.NewSource(time.Now().UnixNano())

	// Use the Source object as basis for a new random number generator
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
