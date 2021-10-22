package main

import "fmt"

func main() {
	fmt.Println(recursiveSum([]int{1, 2, 3, 4, 5}))
}

// summation using divide and conqure - define base case and reduce to it
func recursiveSum(arr []int) int {
	// for empty arrays
	if len(arr) == 0 {
		return 0
	}

	// base case
	if len(arr) == 1 {
		return arr[0]
	}

	// reduction to base case using recursive call
	return arr[0] + recursiveSum(arr[1:])
}
