package main

import (
  "fmt"
)

type Point struct {
  x int32
  y int32
}

func changeX(pt *Point) {
  pt.x = 100
}

type Circle struct {
  radius float64
  center *Point
} 

type Circle2 struct {
  radius float64
  *Point
} 

func main() {
  var p1 Point = Point{1, 2}
  fmt.Println(p1.x)

  p1.x = 3
  fmt.Println(p1.x)

  p2 := Point{1, 2}
  fmt.Println(p2)

  // "y" is gonna be the default int32 value which is 0.
  p3 := Point{x: 3}
  fmt.Println(p3)

  p4 := &Point{y: 3}
  fmt.Println(p4)

  fmt.Println(*p4)
  changeX(p4)
  fmt.Println(*p4)

  p5 := &Point{1, 3}
  fmt.Println(*p5)

  (*p5).x = 5
  fmt.Println(*p5)

  // Same as above with structs we don't have to dereference.
  p5.x = 5
  fmt.Println(*p5)

  p6 := &Point{1, 3}
  c1 := Circle{10, p6}
  fmt.Println(c1)

  c2 := Circle{10, &Point{3, 5}}
  fmt.Println(c2)
  fmt.Println(c1.center.x)

  c3 := Circle2{10, &Point{3, 5}}
  fmt.Println(c3)
  fmt.Println(c3.x)
}

