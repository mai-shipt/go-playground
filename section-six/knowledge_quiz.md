# Knowledge Quiz for Section 6 - Interfaces

#### Q1: When we say that interfaces can be satisfied implicitly, we mean that...
we don't need to write extra code to say that some type satisfies an interface. It doesn't need explicit declarations to establish the relationship (e.g. `implement` or `extend` keywords).

#### Q2: To say that a type satisfies an interface means that...
The type implements all the functions contained in the interface definition.

#### Q3: In the following code, does the square type satisfy the shape interface?
```go
type shape interface {
    area() int
}
 
type square struct {
    sideLength int
}
 
func (s square) area() int {
    return s.sideLength * s.sideLength
}
 
func printArea(s shape) {
    fmt.Println(s.area())
}
```
Yes, because `square` defines the `area` function that returns an int which satisfy its required contract.

#### Q4: Take a look at the following code. Does the rectangle type satisfy the shape interface?
```go
type shape interface {
    area() int
}
 
type square struct {
    sideLength int
}
 
type rectangle struct {
    height float64
    width float64
}
 
func (s square) area() int {
    return s.sideLength * s.sideLength
}
 
func (r rectangle) area() float64 {
    return r.height * r.width
}
 
func printArea(s shape) {
    fmt.Println(s.area())
}
```
No, because `rectangle`'s version of the `area` function returns a float64, but the `shape` interface expects a return type of `int`. Thus, it does not satisfy the required contract.

#### Q5: Take a look at the following code. Type square appears to successfully implement the shape interface, but the implementation of square 's area function looks broken - it always returns a value of 10 no matter what the side length of the square is. Will the shape interface do anything to help us catch this error?
```go
type shape interface {
    area() int
}

type square struct {
    sideLength int
}

func (s square) area() int {
    return 10
}

func printArea(s shape) {
    fmt.Println(s.area())
}
```
No, interfaces are only used to help with types. We can still easily write code that does something completely wrong. The `shape` interface will not do anything to help catch that implementation detail error.

#### Q6: Types that implement the Reader interface are generally used to...
read information from an outside data source into our application.

#### Q7: Imagine that you ask a coworker to create a new type that implements the Reader interface to take data from a text file and print it on the screen. They present you with the following code and they say that this code successfully compiled, so it must be correct. You then review their code and give them feedback. What would you say?
```go
type textFileReader struct {}

func (textFileReader) Read(bs []byte) (int, error) {
    return "Information from a text file"
}
```
I would say that while the `textFileReader` type conforms to the requirements of the `Reader` interface, it doesn't actually implement the desired behavior of reading a file from the hard drive. It instead returns a string directly which isn't correct.
