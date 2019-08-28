package main

import "fmt"

func main() {
	// Case 1
	arr := []int{1, 2, 3}
	fmt.Printf("Case 1: %v\n", sumArray(arr))
}

func sumArray(num []int) int {
	var sum int

	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}
