package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin) 
  fmt.Printf("Type the year you were born: ")
  scanner.Scan()

  // Second argument: Base    Third argument: Size
  // Ff we can't get the text raise an error and store it in the "err".
  // input, err := strconv.ParseInt(scanner.Text(), 10, 64)

  // "_" means ignore the error. 
  input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

  fmt.Printf("You will be %d years old at the end of 2020", 2020-input)
} 
