package main

import "fmt"

func main() {
  var name string = "Kaveh"
  fmt.Println(name)

  var last_name string
  last_name = "Azari"
  last_name = "genius"

  fmt.Println(last_name)

  // error
  // var number uint8 = 260

  var number uint16 = 260
  fmt.Println("Hello World!", number)
}
