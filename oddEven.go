
package main

import "fmt"
import "os"
import "strconv"


func main() {
   num := os.Args[1]
   number, _ := strconv.Atoi(num)
   fmt.Println(num)

   if number % 2 == 0 {
        fmt.Println("Number is even")
    } else {
        fmt.Println("Number is odd")
    }

}
