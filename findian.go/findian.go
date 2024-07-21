package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)

	if containsAll(input, "i", "a", "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

func containsAll(s string, chars ...string) bool {
	for _, char := range chars {
		if !contains(s, char) {
			return false
		}
	}
	return true
}

func contains(s, char string) bool {
	for _, c := range s {
		if string(c) == char {
			return true
		}
	}
	return false
}
