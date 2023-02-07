package main

import (
	"fmt"
	"context"
	"time"
)

// "doSomething" function acts like a function that sends work 
// to one or more goroutines reading from a work channel.
// "doAnother" is the worker and printing the numbers is the work it's doing.
func doSomething(ctx context.Context) {
	ctx, cancelCtx := context.WithCancel(ctx)

	printCh := make(chan int)

	go doAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		printCh <- num
	}

	cancelCtx()

	// With this we give "doAnother" some time to process the canceled context
	// and finish running.
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("doSomething: finished\n")
}

func doAnother(ctx context.Context, printCh <-chan int) {
	for {
    select {
	  	case <-ctx.Done():
			  if err := ctx.Err(); err != nil {
				  fmt.Printf("doAnother err: %s\n", err)
			  }

			  fmt.Printf("doAnother: finished\n")

			  return 
			case num := <-printCh:
				fmt.Printf("doAnother: %d\n", num)
		}	
	}
}

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "myKey", "myValue")

	doSomething(ctx)
}
