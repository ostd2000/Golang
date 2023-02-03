package main

import (
  "fmt"
  "math"
)

// If any of our strucs have the "area" method on them they are a type "Shape".
// "Shape" is a higher level thing that connects these structs.
type Shape interface {
  area() float64
} 

type Circle struct {
  radius float64
}

type Rect struct {
  width float64
  height float64
}

func (r Rect) area() float64 {
  return r.width * r.height
}

func (r *Rect) area2() float64 {
  return r.width * r.height
}

func (c Circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

func (c *Circle) area2() float64 {
  return math.Pi * c.radius * c.radius
}
 
func getArea(s Shape) float64 {
  return s.area()
}

func main() {
  r1 := Rect{3, 5}
  c1 := Circle{3}

  shapes := []Shape{r1, c1}

  for _, shape := range shapes {
    // fmt.Println(shape.area())
    fmt.Println(getArea(shape))
  }

  shapes2 := []Shape{&r1, &c1}

  for _, shape := range shapes2 {
    fmt.Println(getArea(shape))
  }
}
