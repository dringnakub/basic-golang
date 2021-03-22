package main

import (
	"fmt"
)

// Global variable //
var count int = 0

func addCount() {
	count++
}
func main() {
	fmt.Println("Hello DrinkRider")

	// Explicit Declaration //
	var tmp1 int = 0
	tmp1 = 20
	var tmp2 string = "dringrider"
	var tmp3 bool = true

	const tmp4 int = 100
	// cannot assign to tmp4
	// tmp4 = 500

	// Implicit Declaration //
	// var tmp5 int = 0
	tmp5 := 0

	fmt.Println(tmp1)
	fmt.Println(tmp2)
	fmt.Println(tmp3)
	fmt.Println(tmp4)

	fmt.Println(tmp5)

	fmt.Println("count before", count)
	addCount()
	addCount()
	addCount()
	fmt.Println("count after", count)
}
