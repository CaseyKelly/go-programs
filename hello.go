// package clause starts every source file
package main

import (
  "fmt"   // a package in the Go standard library, means format
  "time"  // Package time provides functionality for measuring and displaying time.
  )

func main() {
  fmt.Println("Hello World!")
  t := time.Now()
  fmt.Println(t.Format("Today is Monday, January 2, 2006. The time is 3:04 PM."))
}
