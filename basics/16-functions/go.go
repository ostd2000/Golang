package main

import (
  "fmt"
)

func add(x int, y int) {
  fmt.Println(x + y)
}

// Same as above.
func add2(x, y int) {
  fmt.Println(x + y)
}

func add3(x, y int) int {
  return x + y
}

func add4(x, y int) (int, int) {
  return x + y, x - y
}

// Labeled return values.
func add5(x, y int) (r1 int, r2 int) {
  r1 = x + y
  r2 = x - y
  return
}

// Same as above.
func add6(x, y int) (r1, r2 int) {
  r1 = x + y
  r2 = x - y
  return
}

// defer: defers the execution of the line until we hit the "return" keyword.
func add7(x, y int) (r1, r2 int) {
  defer fmt.Println("We hit the return keyword")

  r1 = x + y
  r2 = x - y

  fmt.Println("before return")

  return
}


func main() {
  add(3, 5)

  ans := add3(3, 5)
  fmt.Println(ans)

  ans2, ans3 := add4(3, 5)
  fmt.Println(ans2, ans3)

  ans4, ans5 := add5(3, 5)
  fmt.Println(ans4, ans5)

  ans6, ans7 := add7(3, 5)
  fmt.Println(ans6, ans7)
}
