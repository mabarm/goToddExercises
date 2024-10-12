package main

import (
	"fmt"
)

func run() {
	fmt.Println("Hi run")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func walk() {
	fmt.Println("Hi walk")
}

func main() {
	fmt.Println("Program started")
	go run()
	walk()
	fmt.Println("Program ended")
}

/* main completed its execution
before run can completely executed
causing bug
*/
