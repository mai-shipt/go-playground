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

# Knowledge Quiz for Section 4 - Value & Reference Types

#### Q1: When we create a slice, Go will automatically create which two data structures?
An array and a structure called "slice descriptor" that records the length of the slice, the capacity of the slice, and a pointer to the actual underlying array.

#### Q2: In the following code snippet, when we pass mySlice to the updateSlice function, is the mySlice value being copied before being passed into the function?
```go
package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "how", "are", "you?"}

	updateSlice(mySlice)

	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
```
Yes, the slice descriptor (the length, capacity, and pointer to the underlying array) are copied, but the original underlying array remains unchanged.

#### Q3: With 'value types' in Go, do we have to worry about pointers if we want to pass a value to a function and modify the original value inside the function?
Yes, because those value values (e.g. integer, float, string, bool, struct) are copied and the original value is unchanged.

#### Q4: With 'reference types' in Go, do we have to worry about pointers if we want to pass a value to a function and modify the original value inside the function?
No, because those reference types (e.g. slices, maps, channels, pointers, functions) directly modifies the original value, so there's no need for pointers.

#### Q5: Is a slice a 'value type' or a 'reference type'
A reference type because a slice contains a reference to the actual underlying list of records.

#### Q6: Here's a tough one. We've been discussing about how to use pointers to avoid passing something to a function by value.  So instead of passing a value to a function, we pass a pointer to that value.  But we've also said many times that Go is a "pass by value" language, which *always* copies arguments that are passed to a function. Take a glance at the following code snippet, which gets a pointer to name called namePointer and prints out the memory address of the pointer itself! Then in the function, we take the pointer that was passed as an argument and print out the address of that pointer too. Do you think the memory address printed by both Println calls will be the same? Why or why not?
```go
package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
```
The log statements will print different addresses because *everything* in Go is pass by value, so when a pointer is passed to a function, a copy is made. But, it is still pointing to the same memory address so changes affect the same underlying value.
