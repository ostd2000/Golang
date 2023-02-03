package main

import "fmt"

func main() {
  // Implicit declaration.
  var number = 5
  fmt.Printf("%T", number)

  // Expression assignment operator.
  number_2 := 5
  fmt.Printf("%T", number_2)

  var number_3 uint64
  fmt.Println(number_3)    // 0

  var bl bool 
  fmt.Println(bl)    // false

  var str string
  fmt.Println(str)    // ""
}
