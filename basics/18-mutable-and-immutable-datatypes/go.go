// "arrays", "slices", and "maps" are mutable and that means they can change. 
// We can literally change some of the values stored in them without creating a new one.

package main

import (
  "fmt"
)

func changeFirst(slice []int) {
  slice[0] = 1000
}

func main() {
  // "int" is immutable.
  var x int = 5
  y := x
  y = 10
  fmt.Println(x, y)    // 5 10
	

  // "slice" is mutable. 
  var x2 []int = []int{1, 2, 3}

  // This means y is equal to the slice that x is pointing to.
  // x and y are pointing to the same underlying variable.
  y2 := x2   
  y2[0] = 100
  fmt.Println(x2, y2)    // [100 2 3] [100 2 3]


  // "map" is mutable.
  var x3 map[string]int = map[string]int{"Hello": 1}
  y3 := x3
  y3["bye"] = -1
  fmt.Println(x3, y3)    // map[Hello:1 bye:-1] map[Hello:1 bye:-1]

  // "array" is mutable: but it behaves differently to the slices and maps.
  var x4 [3]int = [3]int{1, 2, 3}

  // makes a copy of the original array and assigns it to y4.
  y4 := x4

  y4[0] = 100
  fmt.Println(x4, y4)    // [1 2 3] [100 2 3]

  var x5 []int = []int{1, 2, 3}
  fmt.Println(x5)    // [1 2 3]
  changeFirst(x5)
  fmt.Println(x5)    // [1000 2 3]
} 
