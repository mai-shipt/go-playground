# Knowledge Quiz for Section 4 - Pointers

#### Q1: Whenever you pass an integer, float, string, or struct into a function, what does Go do with that argument?
Go is pass by value, so it creates a copy of each argument, and those copies are used inside the function.

#### Q2: What will the following program print out?
```go
package main

import "fmt"
 
func main() {
   name := "Bill"
 
   fmt.Println(&name)
}
```
It will print out the memory address that "Bill" is stored at because it's using the '&' operator.

#### Q3: What is the & operator used for?
It turns a value into a pointer. It retrieves the memory address of a variable which converts that variable's value into a pointer to its location in memory.

#### Q4: When you see a * operator in front of a pointer, what will it turn the pointer into?
It will turn a pointer into a value. The '*' operator dereferences a pointer, allowing access to the actual value stored at that referenced memory address.

#### Q5: When the following program runs, the fmt.Println call reports that the latitude field of newYork is still equal to 40.73. What changes should we make to get the latitude of newYork to update to 41.0?
```go
package main

import "fmt"
 
type location struct {
	longitude float64
    latitude float64
}
 
func main() {
    newYork := location{
        latitude: 40.73,
        longitude: -73.93,
    }
    newYork.changeLatitude()
    fmt.Println(newYork)
}
 
func (lo location) changeLatitude() {
    lo.latitude = 41.0
}
```
Change the receiver type of `changeLatitude` to `*location`, then replace `lo` with `(*lo)` in the function body. This will turn the pointer `lo` into a value type and then update it. See below:
```go
func (lo *location) changeLatitude() {
    (*lo).latitude = 41.0
}
```

#### Q6: Take a look at the following snippet of code. In the 'changeLatitude' function, what is *location in the receiver list (after the word 'func') communicating to us?
```go
package main
 
import "fmt"
 
type location struct {
    longitude float64
    latitude float64
}
 
func main() {
    newYork := location {
        latitude: 40.73,
        longitude: -73.93,
    }
    newYork.changeLatitude()
	fmt.Println(newYork)
}
 
func (lo *location) changeLatitude() {
    (*lo).latitude = 41.0
}
```
It specifies the type of the receiver that the function expects. The `*location` in the receiver list says that the `changeLatitude` function is designed to work with a pointer to a location type (it can modify the original struct).

#### Q7: Take a look at the following program.  What will the Println  function in the main  function print out?
```go
package main
 
import "fmt"
 
func main() {
    name := "Bill"
    updateValue(name)
    fmt.Println(name)
}
 
func updateValue(n string) {
    n = "Alex"
}
```
It will print "Bill". Go is pass by value, so when you pass a variable to a function, a copy is created, leaving the original variable unchanged.

#### Q8: Take a look at the following program.  The changeLatitude function expects a receiver of type pointer to a location struct , but in the main function the receiver is a value type of a struct.  What will happen when this code is executed?
```go
package main
 
import "fmt"
 
type location struct {
    longitude float64
    latitude float64
}
 
func main() {
    newYork := location{
        latitude: 40.73,
        longitude: -73.93,
    }
 
 newYork.changeLatitude()
 
 fmt.Println(newYork)
}
 
func (lo *location) changeLatitude() {
    (*lo).latitude = 41.0
}
```
This program uses a shortcut, where Go will automatically assume that even though `newYork.changeLatitude()` is using a value type we probably meant to pass in a pointer to the `newYork` struct.

#### Q9: Here's a tricky one!  What will the following program print out?
```go
package main
 
import "fmt"
 
func main() {
    name := "Bill"
 
    fmt.Println(*&name)
}
```
It prints out "Bill". Even though the expression `*&name` uses both operators, they "cancel out", returning the original value. The leftmost operator `*` is applied last so the pointer is dereferenced back to the value.
