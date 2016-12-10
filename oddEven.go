// Program to find the given number odd  or even.
package main

import "fmt"
import "os"
import "strconv"

func main() {
   num := os.Args[1]
   number, _ := strconv.Atoi(num)

   if number % 2 == 0 {
        fmt.Println(number, "is even.")
    } else {
        fmt.Println(number, "is odd.")
    }

}
