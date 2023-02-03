package main

import (
  "fmt"
)

func main() {
  val := (true || false) && !false || 5 > 3
  fmt.Printf("%t", val)
} 
