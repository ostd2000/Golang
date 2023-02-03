package main

import (
	"fmt"
)

func main() {
  var x [5]int = [5]int{1, 2, 3, 4, 5}

  // "slice": Copies the whole array.
  var s []int = x[:]
  fmt.Println(s)

  // "slice": Start at 1 go to the end.
  var s2 []int = x[1:]
  fmt.Println(s2)

  // "slice": Start at 1 go to 3 but do not include 3 itself.
  var s3 []int = x[1:3]
  fmt.Println(s3)
  fmt.Println(len(s3))    // 2
  fmt.Println(cap(s3))    // 4

  // Reslicing
  fmt.Println(s3[:cap(s3)])    // [2 3 4 5]

  // Making a slice.
  var a []int = []int{1, 2, 3, 4, 5}
  fmt.Println("len", len(a))    // 5
  fmt.Println("cap", cap(a))    // 5
  a = append(a, 10)
  b := append(a, 10)

  fmt.Println(a)
  fmt.Println(b)

  // Making a slice using "make()".
  c := make([]int, 5)
  fmt.Println(c)
}
