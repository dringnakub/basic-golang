package main

import "fmt"

func main() {
	// var a [5]int
	// a[1] = 10
	// fmt.Println(a)

	var arrayInt = []int{10, 20, 30, 40, 50}
	var arrayString [3]string

	for _, item := range arrayInt {
		fmt.Println("item", item)
	}

	arrayString[0], arrayString[2] = "Dota", "Rov"

	for _, item := range arrayString {
		// if index == 0 && item == "Dota" {
		// 	break
		// }
		fmt.Println("item arrayString", item)
	}
}
