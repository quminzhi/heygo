package main

import (
	"fmt"
)

// Stack entry representing subarray indices
type stackEntry struct {
	left, right int
}

// MergeSortStack performs an iterative (non-recursive) merge sort using a stack.
func MergeSortStack(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	// Stack to store subarray indices
	stack := []stackEntry{{0, n - 1}}

	// Process the stack to divide the array into smaller parts
	var subarrays [][]int
	for len(stack) > 0 {
		// Pop the top element (LIFO order)
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		left, right := top.left, top.right
		if left >= right {
			continue
		}

		mid := left + (right-left)/2
		stack = append(stack, stackEntry{left, mid})      // Push left subarray
		stack = append(stack, stackEntry{mid + 1, right}) // Push right subarray

		// Store subarrays to merge later
		subarrays = append(subarrays, []int{left, mid, right})
	}

	// Merge subarrays in bottom-up order
	for i := len(subarrays) - 1; i >= 0; i-- {
		left, mid, right := subarrays[i][0], subarrays[i][1], subarrays[i][2]
		merge(nums, left, mid, right)
	}
}

// merge function merges two sorted subarrays: nums[left:mid+1] and nums[mid+1:right+1]
func merge(nums []int, left, mid, right int) {
	buf := make([]int, right-left+1)
	l, r, p := left, mid+1, 0

	// Merge two sorted subarrays
	for l <= mid && r <= right {
		if nums[l] <= nums[r] {
			buf[p] = nums[l]
			l++
		} else {
			buf[p] = nums[r]
			r++
		}
		p++
	}

	// Copy remaining elements
	for l <= mid {
		buf[p] = nums[l]
		l++
		p++
	}
	for r <= right {
		buf[p] = nums[r]
		r++
		p++
	}

	// Copy merged result back to original array
	copy(nums[left:right+1], buf)
}

// Example usage
func main() {
	arr := []int{5, 2, 9, 1, 5, 6, 3, 7, 8, 4}
	fmt.Println("Before sorting:", arr)
	MergeSortStack(arr)
	fmt.Println("After sorting:", arr)
}
