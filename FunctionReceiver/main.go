package main

import "fmt"

// Define a Person struct
type Person struct {
	FirstName string
	Age       int
}

// Normal function that accepts a struct by value (won't modify the original)
func updateAgeByValue(p Person, newAge int) {
	p.Age = newAge
	fmt.Println("Inside updateAgeValue:", p.Age)

}

// Normal function that accepts a struct by pointer (modifies the original)
func updateAgeByPointer(p *Person, newAge int) {
	p.Age = newAge
	fmt.Println("Inside updateAgePointer:", p.Age)
}

func main() {
	person1 := Person{FirstName: "John", Age: 25}
	fmt.Println("Before updateAgeByValue:", person1) // Output: {John 25}

	// Pass by value (original struct won't be modified)
	updateAgeByValue(person1, 99)
	fmt.Println("After updateAgeByValue:", person1) // Output: {John 25}

	fmt.Println("--------------")

	fmt.Println("Before updateAgeByPointer:", person1) // Output: {John 25}

	// Pass by pointer (original struct will be modified)
	updateAgeByPointer(&person1, 99)
	fmt.Println("After updateAgeByPointer:", person1) // Output: {John 99}
}

/*
Key Difference Between Functions and Methods
1.With normal functions, you must manually handle how the struct is passed:

1.1 Pass by value: The function works on a copy of the struct.
	Call syntax: updateAgeByValue(person1, 99)
	Function signature: func updateAgeByValue(p Person, newAge int)
1.2 Pass by pointer: The function works on the original struct via a pointer.
	Call syntax: updateAgeByPointer(&person1, 99)
	Function signature: func updateAgeByPointer(p *Person, newAge int)

2. With methods, The call syntax remains the same regardless,
i.e. Go automatically handles whether the struct is passed by pointer or value based on the receiver type (pointer receiver *T or value receiver T).
*/
