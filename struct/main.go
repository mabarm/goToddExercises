package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Employee struct {
	House          int
	Person_Details Person
	Address        // Embedded struct (anonymous field)
}

type Address struct {
	City, Country string
}

func (p Person) fullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) updateAgeMethod(newAge int) {
	p.Age = newAge
}

func updateAge(p *Person, newAge int) {
	p.Age = newAge
	fmt.Println("Inside updateAge:", p.Age)
}

func main() {
	var person1 Person
	fmt.Println(person1)

	person1.Age = 20
	person1.FirstName = "John"
	person1.LastName = "Bog"
	fmt.Println("Person1 is", person1)
	fmt.Println("Person1 is", person1.FirstName)
	fmt.Println("Person1 fullName is", person1.fullName())

	fmt.Println("------------")
	fmt.Println("Person1's new age is before update", person1.Age)
	// person1.updateAgeMethod(102)
	// fmt.Println("Person1's new age is after update", person1.Age)
	updateAge(&person1, 99)
	fmt.Println("Person1's new age is after update", person1.Age)

	person1.FirstName = "NewJohn"
	fmt.Println("Person1 is", person1)
	fmt.Println("Person1 is", person1.FirstName)

	fmt.Println("------------")

	person2 := Person{
		FirstName: "Shreya",
		Age:       30,
		LastName:  "SS",
	}
	fmt.Println("Person2 is", person2, person2.FirstName)

	var employee1 Employee
	employee1.Person_Details.FirstName = "Tim"
	employee1.Person_Details.Age = 40
	fmt.Println("Employee1 is", employee1)
	fmt.Println("------------")

	employee2 := Employee{
		House: 100,
		Person_Details: Person{
			FirstName: "Pierce",
			LastName:  "Mo",
			Age:       30,
		},
		Address: Address{
			City:    "London",
			Country: "UK",
		},
	}
	fmt.Println("Employee2 is before", employee2)
	// updateAge(&employee2.Person_Details, 0)
	// employee2.Person_Details.updateAgeMethod(10)
	fmt.Println("Employee2 is after ", employee2)

	fmt.Printf("Employee2's age is: %d, House is: %d,  City is: %s, Country is: %s, City is: %s \n",
		employee2.Person_Details.Age,
		employee2.House,
		employee2.City,
		employee2.Address.Country,
		employee2.Address.City,
	)

	fmt.Println("------------")

	employee3PD := Person{
		FirstName: "Sheldon",
		LastName:  "Cooper",
		Age:       38,
	}

	employee3 := Employee{
		House:          800,
		Person_Details: employee3PD,
	}
	fmt.Println("Employee3 is ", employee3, employee3.Address, employee3.Person_Details.FirstName)

}

/*
In Go (Golang), a struct is a composite data type that groups together variables under a single name.
These variables, called fields, can have different types,
and a struct is used to model real-world entities or to group related data logically.
*/
