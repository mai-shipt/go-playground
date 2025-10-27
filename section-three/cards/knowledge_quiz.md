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
