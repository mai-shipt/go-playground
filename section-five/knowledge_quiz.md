# Knowledge Quiz for Section 4 - Maps

#### Q1: Can some of the keys in a single map be of type int and others of type string?
No, a map is statistically typed meaning all its keys must be the same type (and same for all its values). Go requires type consistency in a collection like maps.

#### Q2: Can some of the values in a single map be of type int and others of type string?
No, same reasoning as above.

#### Q3: Take a look at the following code.  What would the print statement log out?
```go
package main

import "fmt"
 
func main() {
    m := map[string]string{
        "dog": "bark",
    }
 
    changeMap(m)
    fmt.Println(m)
}
 
func changeMap(m map[string]string) {
    m["cat"] = "purr"
}
```
It would print out: `map[dog: bark cat: purr]` because maps are type references so it will modify the original underlying data structure. Also, since the key "cat" doesn't exist, it will append it to the map.

#### Q4: What would happen if we tried to run the following program? Look closely at all the code in it :)
```go
package main

import "fmt"
 
func main() {
    m := map[string]string{
        "dog": "bark",
        "cat": "purr",
    }
 
    for key, value := range m {
        fmt.Println(value)
    }
}
```
The compiler would throw an error because the variable `key` was created but never used. Otherwise, if it uses a blank identifier `_` for `key` then it would successfully execute and print out: 
```
bark
purr
```
