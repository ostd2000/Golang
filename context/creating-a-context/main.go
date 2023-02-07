package main

import (
	"fmt"
	"context"
)

// It is recommended to put the "context.Context" parameter 
// as the first parameter in a function.
func doSomething(ctx context.Context) {
	fmt.Println("Doing something")
}

func main() {
	// Creates an empty context or a starting context.
	ctx := context.TODO()

	// Another way to create an empty context.
	// but it is designed to be used where you intend to start a known context.
	ctx2 := context.Background()

	doSomething(ctx)
	doSomething(ctx2)
}
