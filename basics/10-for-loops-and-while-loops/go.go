package main

import (
  "fmt"
)

func main() {
  x := 3
  for x < 5 {
    fmt.Println(x)
    x++
  }

  fmt.Println("----------")

  for x := 3; x < 5; x++ {
    fmt.Println(x)
  }

  fmt.Println("----------")

  // We have the concept of "continue" and "break" in Golang.
  for x := 0; x <= 100; x++ {
    if x != 0 && x % 3 == 0 && x % 5 == 0 {
      fmt.Println(x)

      // break
      // continue
    }
  } 
}
