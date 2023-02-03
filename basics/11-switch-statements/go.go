package main

import (
  "fmt"
)

func main() {
  ans := 10

  switch ans {
    case 1, -1:
      fmt.Println("It's one or negative one")
    case 2:
      fmt.Println("It's two")
    default:
      fmt.Println("not a case")
  } 

  switch {
    case ans > 0:
      fmt.Println("It's greater than zero")
    case ans < 0:
      fmt.Println("It's less than zero")
    default:
      fmt.Println("Zero")
  }
}
