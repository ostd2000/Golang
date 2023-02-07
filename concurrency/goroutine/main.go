// The "WaitGroup" primitive works by counting how many things
// it needs to wait for using the "Add", "Done" and "Wait" functions. 
// The "Add" function increases the count by the number provided to the function,
// and "Done" decreases the cout by one. 
// The "Wait" function can then be used to wait until the count reaches zero, 
// meaning that "Done" has been called enough times to offset the calls to "ADD".
// Once the count reaches zero, the "Wait" function will return 
// and the program will continue running.

package main

import (
	"fmt"
	"sync"
)

func generateNumbers(total int, wg *sync.WaitGroup) {
	// "defer" waits for the function to finish running,
	// then executes what's specified.
	defer wg.Done()

	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Generating number %d\n", idx)
	}
}

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()

	for idx := 1; idx <= 3; idx++ {
		fmt.Printf("Printing number %d\n", idx)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	
	go printNumbers(&wg)
	go generateNumbers(3, &wg)

	fmt.Println("Waiting for goroutines to finish...")
  
	// When you remove the call to "Wait", the program will no longer wait for
	// both functions to finish running before it continues.
	wg.Wait()

	fmt.Println("Done!")
}
