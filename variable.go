
package main

import "fmt"

type temperature struct{
  Value float64
  Type string
}
func main() {
  var a int = 10
  var b int = 20
  var name string = "Hello"
  var user struct {
    Name string
    Email string
  }
  temp:= temperature{Value:5.6, Type:"celcius"}


  fmt.Printf("%#v, %#v, %#v\n", a, b, name)
  fmt.Printf("%#v\n",user)
  fmt.Printf("%#v\n",temp)
}
