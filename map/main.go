package main

import "fmt"

func main() {
	studentAge := make(map[string]int)

	studentAge["Bob"] = 10
	fmt.Printf("Bob's age is %d\n", studentAge["Bob"])

	studentAge["Banard"] = 20

	for name, value := range studentAge {
		fmt.Printf("%s is of %d years old\n", name, value)
	}

	fmt.Printf("%d\n", studentAge["Sham"])

	name, exists := studentAge["Sham"]
	fmt.Printf("Does Sham exists in map %t with value %d \n", exists, name)

	delete(studentAge, "Bob")

	nameBob, existsBob := studentAge["Bob"]

	fmt.Printf("Does Bob exist in map: %t with value %d\n", existsBob, nameBob) //value is 0 although key doesn't exist

	//Direct initialisation
	teacherAge := map[string]int{
		"Ren":  50,
		"Mary": 40,
	}

	for name, age := range teacherAge {
		fmt.Printf("%s age is %d\n", name, age)
	}

}
