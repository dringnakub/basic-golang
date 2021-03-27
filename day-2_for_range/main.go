package main

import "fmt"

func main() {
	games := []string{"ROV", "Dota", "Cabal Mobile"}
	// _ ,  ignore index //

	for index, item := range games {
		// condition //
		if item == "ROV" {
			fmt.Println("Woow ")
		}
		fmt.Printf("%d. %s\n", index+1, item)
	}
}
