package main

import "fmt"

// Type of 'deck' has all features of a slice of strings
type deck []string

// print() method is related to "type" deck
// Any variable of type "deck" have access to this "print" method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
