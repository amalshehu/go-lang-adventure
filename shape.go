package main

import "fmt"
import "os"
// import "strconv"

const Pi = 3.14
type Circle struct {
    radius  float64
}

type Rectangle struct {
    length, width  float64
}

func main() {
  input := os.Args[1]

  func (c Circle) float64 {
    return Pi*c.radius*c.radius
  }
  func (r Rectangle) int {
    return r.length*r.width
  }

}
