package main

import "fmt"

func main() {
	value := 20
	var valuePointer *int = &value

	fmt.Println("value is ", value)
	// Address //
	fmt.Println("valuePointer is ", valuePointer)
	// get value //
	fmt.Println("valuePointer is ", *valuePointer)

	// pass address //
	changeMessage(&value, 400)
	fmt.Println("changeMessage call 1", value)

	// indirect //
	changeMessage(valuePointer, 500)
	fmt.Println("changeMessage call 2", value)

	changeMessage(valuePointer, 900)
	fmt.Println("changeMessage call 3", *valuePointer)
}

func changeMessage(aPointer *int, newInt int) {
	// change value //
	*aPointer = newInt
}
