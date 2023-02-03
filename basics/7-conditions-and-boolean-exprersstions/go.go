package main

import (
  "fmt"
)

func main() {
  a := 3
  b := 5
  c := 5.5
  d := "kaveh"
  e := "Kaveh"
  f := 'a'
  g := 'b'

  fmt.Println(a < b)
  fmt.Println(a == b)
  fmt.Println(float64(a) == c)
  fmt.Println(float64(b)+1.5 != c)
  fmt.Println(d == e)
  fmt.Println(e < d)
  fmt.Println(f == g)
}
