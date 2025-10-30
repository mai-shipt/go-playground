package main

import "fmt"

// Creates a new type 'bot' interface where if you have a fxn called
// 'getGreeting' and you return a 'string' then you are automatically
// a member of type 'bot' (no need to use 'extend' or 'implement'
// keywords like other languages)
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// Just needed one printGreeting() fxn with one implementation
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
