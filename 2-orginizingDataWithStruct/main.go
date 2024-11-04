package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// alex := person{}
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// alex.contact.email = "alex@gmail.com"
	// alex.contact.zipCode = 12345
	// fmt.Println(alex.contact.email)

	// Create a new person struct for Jim
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}

	// // Get the memory address of jim using & operator and store in jimPointer
	// jimPointer := &jim

	// // Call updateName method on the pointer to update jim's first name to "jimmy"
	// jimPointer.updateName("jimmy")

	// // Print the updated jim struct
	// jim.print()
	jim.updateName("jimmy")
	jim.print()
}

// updateName takes a pointer to person and updates the firstName field
// pointerToPerson is the receiver parameter that is a pointer to person struct
// newFirstName is the new name to update to
func (pointerToPerson *person) updateName(newFirstName string) {
	// Dereference the pointer using * operator to update the actual value
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p)

	fmt.Printf("%+v", p)
}
