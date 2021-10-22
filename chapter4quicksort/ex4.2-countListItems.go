package main

import "fmt"

func main() {
	fmt.Println(recursiveCountListItems([]string{"apple", "banana", "lemon"}))
}

// count items of list using recursion
func recursiveCountListItems(arr []string) int {
	if len(arr) == 0 {
		return 0
	}

	// increment count each time moving one index
	return 1 + recursiveCountListItems(arr[1:])
}
