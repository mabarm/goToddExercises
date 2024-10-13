//Pointer Receiver vs. Value Receiver in Methods

package main

import "fmt"

type Person struct {
	FirstName string
	Age       int
}

// Method with a pointer receiver (modifies the original struct)
func (p *Person) updateAgePointer(newAge int) {
	p.Age = newAge
	fmt.Println("Inside updateAgePointer:", p.Age)
}

// Method with a value receiver (modifies a copy, not the original)
func (p Person) updateAgeValue(newAge int) {
	p.Age = newAge
	fmt.Println("Inside updateAgeValue:", p.Age)

}

func main() {
	person1 := Person{FirstName: "John", Age: 25}
	fmt.Println("Before updateAgePointer:", person1) // Output: {John 25}

	// Call the method (Go automatically handles the pointer)
	person1.updateAgePointer(99)

	// The original struct is modified
	fmt.Println("After updateAgePointer:", person1) // Output: {John 99}

	fmt.Println("--------------")

	person2 := Person{FirstName: "John", Age: 25}
	fmt.Println("Before updateAgeValue:", person2) // Output: {John 25}

	// Call the method (Go passes the value, i.e., a copy of person1)
	person2.updateAgeValue(99)

	// The original struct is NOT modified
	fmt.Println("After updateAgeValue:", person2) // Output: {John 25}
}

/*
Go automatically handles pointer dereferencing or struct copying depending on the receiver type

How Go Handles It Internally
1. If the method has a pointer receiver (e.g., func (p *Person)),
	Go will automatically take the address of the struct when calling the method on a struct value (person1.updateAgePointer(99) is equivalent to (&person1).updateAgePointer(99)).
2. If the method has a value receiver (e.g., func (p Person)),
	Go will automatically create a copy of the struct when calling the method (person1.updateAgeValue(99) passes a copy of person1).
*/
