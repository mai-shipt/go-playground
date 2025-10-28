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

	mai.print()
	mai.updateName("Mayo")
	mai.print()
}

// print() prints the person's information to standard out
// This function is tied to the person struct by using a value receiver
func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// updateName() SHOULD change the given person's first name, but why is it not? Fix TBD
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// Leaving this original 'main' commented out for posterity with notes
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
