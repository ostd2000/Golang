package main

import (
  "fmt"
)

func main() {
  var arr[3] int
  fmt.Println(arr)    // [0 0 0]

  var arr2[3] string   
  fmt.Println(arr2)    // [ , , ]

  var arr3[3] int 
  arr3[0] = 100
  fmt.Println(arr3)
  fmt.Println(arr3[0])

  arr4 := [3]int{1, 2, 3}
  fmt.Println(arr4)
  fmt.Println(len(arr4))

  sum := 0 
  for i := 0; i < len(arr); i++ {
    sum += arr[i]
  }

  fmt.Println(sum)

  arr2D := [2][3]int{{1, 2, 3}, {1, 2, 3}}
  fmt.Println(arr2D[0][2]) 
} 
