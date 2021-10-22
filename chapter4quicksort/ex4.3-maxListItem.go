package main

import "fmt"

func main() {
	fmt.Println(maxListItem([]int{1, 12, 3, 8, 4}))
}

// find max val in list using divide and conqure
func maxListItem(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	var max = arr[0]
	if max > arr[1] {
		arr[1] = max
	}

	// using recursive call
	return maxListItem(arr[1:])
}
