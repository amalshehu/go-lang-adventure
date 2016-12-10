
package main

import "fmt"


func main() {
  var bag =  map[string]int{
    "Pen" : 2,
    "paper" : 4,
    "Charts" : 8,
  }
  for key, value:= range bag{
    fmt.Printf(key, value)
  }
}
