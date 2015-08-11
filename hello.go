// package clause starts every source file
package main

import (
	// a package in the Go standard library, means format
	"fmt"

	// Package math provides basic constants and mathematical functions.
	"math/rand"

	// Package time provides functionality for measuring and displaying time. Package rand implements pseudo-random number generators.
	"time"
)

func main() {
	fmt.Println("Hello World!")

	t := time.Now()
	fmt.Println(t.Format("Today is Monday, January 2, 2006. The time is 3:04 PM."))

	rand.Seed(t.Unix())
	fmt.Println(rand.Intn(10), "is a random number")
}
