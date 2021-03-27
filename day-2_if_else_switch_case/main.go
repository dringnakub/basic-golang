package main

import (
	"fmt"
	"time"
)

func main() {
	// var someValue int = 22
	if someValue := 22; someValue == 20 {
		fmt.Printf("value is %d\n", someValue)
	} else {
		fmt.Printf("value not %d\n", someValue)
	}

	valueTest := 24
	if valueTest == 23 {
		fmt.Println(valueTest, "is negative")
	} else if valueTest == 25 {
		fmt.Println(valueTest, "wooow")
	} else {
		fmt.Println(valueTest, "more")
	}

	i := 2
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("default")
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
