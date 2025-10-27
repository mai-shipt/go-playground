# Knowledge Quiz for Section 3 - Functions and Return Types

#### Q1: Which of the following represents a slice where each element in it is of type int?
* A: []string{}
* B: []int{}
* C: []float{}

B correctly represents a slice in Go where each element is of type int.

#### Q2: Is the following code valid?
```go
colors := []strings{"Red", "Yellow", "Blue"}
```
No, because "strings" (plural) is not a valid type! It should be "string" (singular)

#### Q3: How do we iterate through each element in a slice and print out its value?
By using the `for... range` syntax in Go to iterate through each element in a slice, which allows access to the index and value. Like below:
```go
colors := []string{"Red", "Yellow", "Blue"}
for index, color := range colors {
	fmt.Println(index, color)
}
```

#### Q4: Would the following code compile successfully?  Try it out yourself!
```go 
for index, card := range cards {
    fmt.Println(card)
}
```
No, because every variable we declare must be used in our code. In this case, 'index' is not being used. We should use a '_' blank identifier to ignore the 'index'.

#### Q5: Can a slice have both values of type 'int' and of type 'string' in it?
No, because a slice can only have one type of value in it. In Go, slices are homogenerous data structures, meaning they can only contain elements of the same type to ensure type safety and consistency in data handling.
