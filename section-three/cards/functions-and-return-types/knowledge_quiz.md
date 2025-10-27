# Knowledge Quiz for Section 3 - Functions and Return Types

#### Q1: If a function returns a value, do we have to annotate, or mark, the function with the type of value that is being returned?
Yes, every function that returns a value must indicate what type of value it's returning.

#### Q2: The Go compiler is presenting an error message about the following function.  What should we do to fix the error?
```go
func getSize() {
    return "30 meters"
}
```
Add a return type of 'string' to the function, like "func getSize() string {"

#### Q3: Is the following function valid?
```go
func estimatePi() float64 {
    return 3.14
}
```
Yes, it defines a function that returns a float64 value of 3.14. 

#### Q4: Is the following code valid?  Why or why not?
```go 
package main

import "fmt"

func main() {
	fmt.Println(getTitle())
}

func getTitle() string {
	return "Harry Potter"
}
```
Yes, it is valid because functions can be called in any order in Go as long as they are defined within the same package.

#### Q5: Here's a tough one!  Imagine we have two files in the same folder with the same package name.  Will the following code successfully compile?  This might take a little experimentation on your side.  If you do try this out yourself, run your code with the command go run main.go state.go .
In main.go:
```go
package main

func main() {
	printState()
}
```

In a separate file called state.go:
```go
package main

import "fmt"

func printState() {
	fmt.Println("California")
}
```
Yes, because both files are a part of the same package. Files in the same package can freely call functions defined in other files.
