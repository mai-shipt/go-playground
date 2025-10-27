# Knowledge Quiz for Section 3 - Variable Declarations

#### Q1: Is the following a valid way of initializing and assigning a value to a variable?
```go
var bookTitle string = "Harry Potter"
```
Yes

#### Q2: Is the following a valid way of initializing and assigning a value to a variable?
```go
fruitCount := 5
```
Yes, it is using the shorthand ':=' variable declaration. But, it must be inside a function (not global or at the package level).

#### Q3: After running the following code, Go will assume that the variable quizQuestionCount is of what type?
```go
quizQuestionCount := 10
```
Integer

#### Q4: Will the following code compile?  Why or why not?
```go 
paperColor := "Green"
paperColor := "Blue"
```
No, because a variable can only be initialized one time in the same scope. In this case, the ':=' operator is being used to initialize 'paperColor' two times. It will lead to a compilation error.

#### Q5: Are the two lines following ways of initializing the variable 'pi' equivalent?
```go
pi := 3.14
var pi float64 = 3.14
```
Yes

#### Q6: Is the following code valid? (Use the Go Playground at https://play.golang.org/ to quickly run a snippet of code)
```go
package main

import "fmt"

deckSize := 20

func main() {
	fmt.Println(deckSize)
}
```
No, because you cannot use the short variable declaration ':=' at the package level.

#### Q7: We might be able to initialize a variable and then later assign a variable to it.  Is the following code valid?
```go
package main

import "fmt"

func main() {
	var deckSize int
	deckSize = 52
	fmt.Println(deckSize)
}
```
Yes, we can initialize and then later assign a value to a variable.

#### Q8: Is the following code valid?
```go 
package main

import "fmt"

var deckSize int

func main() {
	deckSize = 50
	fmt.Println(deckSize)
}
```
Yes, we can initialize a variable outside a function, we just can't assign a value to it using the := syntax.

#### Q9: Is the following code valid?  Why or why not?
```go 
package main

import "fmt"

func main() {
	deckSize = 52
	fmt.Println(deckSize)
}
```
No, because 'deckSize' is assigned a value before it is initialized. It must first be initialized with the ':=' operator or the 'var variableName type' syntax.
