package main

import "fmt"

func main() {
	fnFor()
	fnWhile()
	fnWhileWithBreak()
}

func fnFor() {
	for index := 0; index < 10; index++ {
		fmt.Printf("For Index %d\n", index)
	}
}
func fnWhile() {
	index := 0
	for index < 5 {
		index++
		fmt.Printf("While Index %d\n", index)
	}
}

func fnWhileWithBreak() {
	index := 1
	for {
		if index > 5 {
			break
		}
		fmt.Printf("While-Index break %d\n", index)
		index++
	}
}
