package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker started %d \n", i)

	fmt.Printf("worker ended %d \n", i)
}
func main() {
	fmt.Println("Main started")

	var wg sync.WaitGroup

	for i := 0; i <= 2; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("Main ended")

}
