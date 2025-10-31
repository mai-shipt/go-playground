# Knowledge Quiz for Section 7 - Channels and Go Routines

#### Q1: Go Routines and Channels are tough, so let's start with the basics! Which of the following best describes what a go routine is?
A separate line of code execution that can be used to handle blocking code. It provides a way to run functions concurrently.

#### Q2: What's the purpose of a channel?
For communications and synchronization between go routines, allowing them to send and receive messages seamlessly.

#### Q3: Take a look at the following program. Are there any issues with it?
```go
package main

import (
	"fmt"
)

func main() {
	greeting := "Hi There!"

	go (func() {
		fmt.Println(greeting)
	})()
}
```
The `greeting` variable is referenced from directly in the go routine, which might lead to issues if we eventually start to change the value of `greeting`. The program will also likely exit before the `fmt.Println` function has a chance to print anything out to the terminal. 

#### Q4: Here's a tough one - Is there any issue with the following code?
```go
package main

func main() {
	c := make(chan string)
	c <- []byte("Hi there!")
}
```
Yes, because the channel is expecting values of type `string` but we are passing in a value of type byte slice which not the same thing as a `string`.

#### Q5: Another tough one! Is there any issue with the following code?
```go
package main

func main() {
	c := make(chan string)
	c <- "Hi there!"
}
```
The syntax is valid, but the program will never exit because it will wait for something to receive the value we're passing into the channel.

#### Q6: Ignoring whether or not the program will exit correctly, are the following two code snippets equivalent?
Snippet #1: 
```go
package main
 
import "fmt"
 
func main() {
    c := make(chan string)
    for i := 0; i < 4; i++ {
        go printString("Hello there!", c)
    }
 
 for s := range c {
     fmt.Println(s)
 }
}
 
func printString(s string, c chan string) {
    fmt.Println(s)
    c <- "Done printing." 
}
```
Snippet #2:
```go
package main
 
import "fmt"
 
func main() {
	c := make(chan string)
 
    for i := 0; i < 4; i++ {
        go printString("Hello there!", c)
    }
 
    for {
        fmt.Println(<- c)
    }
}

func printString(s string, c chan string) {
    fmt.Println(s)
    c <- "Done printing." 
}
```
They are the same: they spawn 4 goroutines to print a string and send a completion signal to the channel. Although they use different mechanisms for reading from the channel, they both effectively achieve the same result which is to print messages until all goroutines finish.
