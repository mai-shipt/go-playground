package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// Define a struct
// This is a custom type called 'person' with these fields/properties
type person struct {
	firstName   string
	lastName    string
	contactInfo // Anonymous field used as or "promoted" as field name too
}

func main() {
	mai := person{
		firstName: "Mai",
		lastName:  "Thao",
		contactInfo: contactInfo{
			email:   "maithao@demo.com",
			zipCode: 12345,
		},
	}
	fmt.Printf("%+v\n", mai)
}

//func main() {
//	// Declare a struct
//	// Create a new person with these values
//	mai := person{firstName: "Mai", lastName: "Thao"}
//	fmt.Println(mai)
//
//	// Another way to declare a struct
//	var alex person
//	fmt.Println(alex)
//	fmt.Printf("%+v", alex) // Print all different field names and values
//	alex.firstName = "Alex"
//	alex.lastName = "Anderson"
//	fmt.Printf("%+v", alex)
//}
