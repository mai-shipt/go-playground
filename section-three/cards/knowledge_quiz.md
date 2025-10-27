# Knowledge Quiz for Section 3 - Functions with Receivers

#### Q1: What would the following code print out?
```go
package main
 
import "fmt"
 
type book string
 
func (b book) printTitle() {
    fmt.Println(b)
}
 
func main() {
    var b book = "Harry Potter"
    b.printTitle()
}
```
This is valid Go code that prints out "Harry Potter". It defines a 'book' type with a 'printTitle()' method that prints its title, which is set to "Harry Potter".

#### Q2: Complete the sentence! "By creating a new type with a function that has a receiver, we..."
are adding a 'method' to any value of that type. Because in Go, when you create a new type with a function that has a receiver, that method can be called on instances of that type.

#### Q3: In the following snippet, what does the variable 'ls' represent?
```go
type laptopSize float

func (ls laptopSize) getSizeOfLaptop() {
    return ls
}
```
The 'ls' variable is a value of the type 'laptopSize'.

#### Q4: Is the following code valid?
```go 
type laptopSize float64

func (this laptopSize) getSizeOfLaptop() laptopSize {
    return this
}
```
Yes, this is valid and will compile, but it breaks Go convention because Go avoids any mention of 'this' or 'self'. We don't ever reference a receiver value as 'this' or 'self', just a short 1-3 letter.

# Knowledge Quiz for Section 3 - Multiple Return Values

#### Q1: In the following code snippet, what will the value and type of 'title' and 'pages' be?
```go
func getBookInfo() (string, int) {
    return "War and Peace", 1000
}

title, pages := getBookInfo()
```
Title will be a string with value "War and Peace" and pages will be an int with value 1000.

#### Q2: What will the following program log out?
```go
package main
 
import "fmt"
 
func main() {
    color1, color2, color3 := colors()
 
    fmt.Println(color1, color2, color3)
}
 
func colors() (string, string, string) {
    return "red", "yellow", "blue"
}
```
"red" "yellow" "blue"

#### Q3: What will the following program log out?
```go
package main

import "fmt"

func main() {
	c := color("Red")

	fmt.Println(c.describe("is an awesome color"))
}

type color string

func (c color) describe(description string) (string) {
	return string(c) + " " + description
}
```
"Red is an awesome color" because the program constructs a string by concatenating the color "Red" with the description "is an awesome color" and then logs the complete string to the console.

`string(c)` is needed in the return statement to explicitly cast it to a string (it is a custom type `color`, not the primitive base type `string`)

#### Q4: Which of the following best explains the describe  function listed below?
```go 
package main

import "fmt"

func main() {
	c := color("Red")

	fmt.Println(c.describe("is an awesome color"))
}

type color string

func (c color) describe(description string) (string) {
	return string(c) + " " + description
}
```
`describe` is a function with a value receiver of type `color` that requires an argument of type `string`, then returns a value of type `string`

#### Q5: This is a tricky question about something that we'll cover in much greater detail later on.  Think back to our current "cards" program, where we have the following code.
```go 
func main() {
    cards := newDeck()
    
    hand, remainingCards := deal(cards, 5)
    
    hand.print()
    remainingCards.print()
}
```
After calling "deal" and passing in "cards", does the list of strings that the "cards" variable point at change?  In other words, did we modify the 'cards' slice by calling 'deal'?

No, `cards` will be the same before and after calling `deal` because we created two new references that point at subsections of the `cards` slice. We never directly modified the slice that `cards` is pointing at (aka the original).
