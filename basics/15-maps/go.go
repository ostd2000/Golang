package main

import (
  "fmt"
)

func main() {
  var mp map[string]int = map[string]int{
    "apple": 1,
    "pear": 2,
    "orange": 9,
  }

  fmt.Println(mp["apple"])

  mp["apple"] = 10

  delete(mp, "pear")

  fmt.Println(mp)

  // Empty "map".
  // mp2 := make(map[string]int)

  // Checking if a key exist in a map: 
  val, ok := mp["apple"]
  fmt.Println(val, ok)    // 10 true

  val2, ok2 := mp["Kaveh"]
  fmt.Println(val2, ok2)    // 0 false

  // Getting the length of a map.
  fmt.Println(len(mp))
}
