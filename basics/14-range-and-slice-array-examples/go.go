package main

import (
  "fmt"
)

func main() {
  var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

  // for i := 0; i < len(a); i++ {
  // 	fmt.Println(a[i])
  // }

  for i, element := range a {
    fmt.Printf("%d: %d   ", i, element)
  }

  fmt.Println()

  // "_": Anonymous variable;
  // if we don't wanna use "i":
  for _, element := range a {
    fmt.Printf("%d ", element)
  }

  fmt.Println()

  var a2 []int = []int{1, 2, 3, 4, 1, 2}
  for i, element := range a2 {
    for j := i + 1; j < len(a2); j++ {
      if element == a2[j] {
	fmt.Printf("%d ", element)
      }
    }
  }
}
