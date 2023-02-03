package main

import "fmt"

type Student struct {
  name   string
  grades []int
  age    int
}

func (s *Student) setAge(age int) {
  s.age = age
}

func (s Student) getAverageGrade() float32 {
  sum := 0
  for _, v := range s.grades {
    sum += v
  }

  return float32(sum) / float32(len(s.grades))
}

func (s Student) getMaxGrade() int {
  curMax := 0
  for _, v := range s.grades {
    if v > curMax {
      curMax = v
    }
  }

  return curMax
}

func main() {
  s1 := Student{"Kaveh", []int{10, 20, 30, 40, 50}, 21}
  fmt.Println(s1)
  s1.setAge(18)
  fmt.Println(s1)

  average := s1.getAverageGrade()
  fmt.Println(average)

  max := s1.getMaxGrade()
  fmt.Println(max)
}
