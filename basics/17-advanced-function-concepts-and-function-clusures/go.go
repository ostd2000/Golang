package main

import (
  "fmt"
)

func test(x int) {
  fmt.Println("hello!", x)
}

func test4(myFunc func(int) int) {
  fmt.Println(myFunc(5))
}

func returnFunc(s string) func() {
  // This is a function closure because it uses a value that we've defined outside of the function, inside of it.   
  return func() { fmt.Println(s) }
}

func main() {
  x := test
  x(5)    // same as: test()

  test2 := func() {
    fmt.Println("hello!")
  }

  test2()

  test3 := func(x int) int {
    return x * -1
  }(10)

  fmt.Println(test3)

  test5 := func(x int) int {
    return x * -1
  }

  test6 := func(x int) int {
    return x * 5
  }

  test4(test5)
  test4(test6)

  func() {
    fmt.Println("call immediately!")
  }()

  returnFunc("Hello!")()
}
