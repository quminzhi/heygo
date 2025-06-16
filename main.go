package main

import (
	"fmt"
	"heygo/bsearch"
	"heygo/sort"
)

// Example usage
func main() {
	arr := []int{5, 2, 9, 1, 5, 6, 3, 7, 8, 4}
	fmt.Println("Before sorting:", arr)
	sort.MergeSort(arr, 0, len(arr)-1)
	fmt.Println("After sorting:", arr)
	bsearch.SearchRange(arr, 3)
}
