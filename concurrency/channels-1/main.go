package main

import (
	"fmt"
	"sync"
)

// "ch chan<- int" means write-only.
// If we don't specify the arrow the channel will have permission
// for both reading and writing.
func generateNumbers(total int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Sending %d to channel\n", idx)

		ch <- idx
	}
}

// "ch <-chan int" means read-only.
func printNumbers(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("Read %d from channel\n", num)
	}
}

func main() {
		var wg sync.WaitGroup 

		numberChan := make(chan int)

		wg.Add(2)

		go printNumbers(numberChan, &wg)

		generateNumbers(3, numberChan, &wg)

		// "close" causes the "for ... range" loop in the "printNumberr" to exit.
		// Since using "range" to read from a channel contnues 
		// until the channel it's reading from is closed.
		close(numberChan)

		fmt.Println("Waiting for goroutines to finish...")
    
		wg.Wait()

		fmt.Println("Done!")
}
