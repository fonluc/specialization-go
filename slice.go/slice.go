package main

import (
	"fmt"
)

func main() {
	var nums []int

	for i := 0; i < 4; i++ {
		var n int
		fmt.Print("Enter an integer: ")
		fmt.Scan(&n)
		nums = append(nums, n)
		nums = bubbleSort(nums)
		fmt.Println("Sorted slice:", nums)
	}
}

func bubbleSort(slice []int) []int {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}
