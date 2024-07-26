package main

import "fmt"

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	var n int

	fmt.Println("Digite 10 inteiros positivos:")
	positiveNumbers := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&n)
		positiveNumbers[i] = n
	}

	BubbleSort(positiveNumbers)
	fmt.Println("Sequência ordenada (positivos):", positiveNumbers)

	fmt.Println("Digite 10 inteiros (com pelo menos um negativo):")
	mixedNumbers := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&n)
		mixedNumbers[i] = n
	}

	BubbleSort(mixedNumbers)
	fmt.Println("Sequência ordenada (mista):", mixedNumbers)
}
