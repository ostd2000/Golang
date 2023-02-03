// When "*" comes before a data type like this: "*string" it means simply pointer
// but when it comes before a variable like this: "*str" it means dereferecing.

package main

import (
  "fmt"
)

// "*string" means i want the pointer not the value.
func changeValue(str *string) {
  *str = "changed!"
}

func changeValue2(str string) {
  str = "changed!"
}

func main() {
  x := 5

  // "&x" gives us the reference of x.
  fmt.Println(&x)    // 0xc000016098

  y := &x
  fmt.Println(x, y)    // 5 0xc000016098
	
  // dereferencing the pointer.
  *y = 10
  fmt.Println(x, y)    // 10 0xc000016098

  toChange := "hello"
  fmt.Println(toChange)    // hello
  changeValue(&toChange)
  fmt.Println(toChange)    // changed!

  toChange2 := "hello"
  fmt.Println(toChange2)    // hello
  changeValue2(toChange2)
  fmt.Println(toChange2)    // hello


  toChange3 := "hello"
  var pointer *string = &toChange3
  fmt.Println(pointer)    // 0xc000088260
  fmt.Println(*pointer)    // hello
}
