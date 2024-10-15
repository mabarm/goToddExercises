package main

import "fmt"

func changeValue(a int) {
	a = 20
	fmt.Printf("Value of a in changeValue  is %v\n", a)
}

func changeValuePointer(a *int) {
	*a = 15
	fmt.Printf("Value of a in changeValuePointer is %v\n", *a)
}

func main() {
	var a int = 10
	var b *int
	b = &a
	fmt.Printf("Value of a is %v and b is %v \n", a, b)
	fmt.Printf("Type of a is %T and b is %T \n", a, b)

	*b = 99
	fmt.Printf("Value of a is %v and b is %v \n", a, b)

	changeValue(a)
	fmt.Printf("Value of a is %v and b is %v \n", a, b)

	changeValuePointer(&a)
	fmt.Printf("Value of a is %v\n", a)

}
