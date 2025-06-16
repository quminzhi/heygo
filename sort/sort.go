package sort

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

// QuickSort sort an array of numbers in [left, right]
func QuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	// Create a new random number generator with a custom source
	rg := rand.New(rand.NewSource(time.Now().UnixNano()))
	pivotIndex := rg.Intn(right-left) + left
	pivot := nums[pivotIndex]

	l, r := left, right
	for l <= r {
		for nums[l] < pivot {
			l++
		}
		for nums[r] > pivot {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	QuickSort(nums, left, r)
	QuickSort(nums, l, right)
}

// QuickSortNR is a non-recursive implementation with stack
func QuickSortNR(nums []int) {
	if len(nums) <= 1 {
		return
	}

	stack := []struct{ left, right int }{{0, len(nums) - 1}}
	rg := rand.New(rand.NewSource(time.Now().UnixNano()))

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left, right := top.left, top.right

		if left >= right {
			continue
		}

		// Random pivot selection
		pivotIndex := rg.Intn(right-left+1) + left
		pivot := nums[pivotIndex]

		// Partitioning
		l, r := left, right
		for l <= r {
			for nums[l] < pivot {
				l++
			}
			for nums[r] > pivot {
				r--
			}
			if l <= r {
				nums[l], nums[r] = nums[r], nums[l]
				l++
				r--
			}
		}

		// Push subarrays onto the stack
		if left < r {
			stack = append(stack, struct{ left, right int }{left, r})
		}
		if l < right {
			stack = append(stack, struct{ left, right int }{l, right})
		}
	}
}

// Parallel threshold (minimum size before using goroutines)
const threshold = 1000

// ParallelQuickSort sorts a nums in parallel
func ParallelQuickSort(nums []int) {
	var wg sync.WaitGroup
	parallelQuickSort(nums, 0, len(nums)-1, &wg)
	wg.Wait() // Wait for all goroutines to finish
}

// Helper function for parallel QuickSort
func parallelQuickSort(nums []int, left, right int, wg *sync.WaitGroup) {
	if left >= right {
		return
	}

	// Random pivot selection
	rg := rand.New(rand.NewSource(time.Now().UnixNano()))
	pivotIndex := rg.Intn(right-left+1) + left
	pivot := nums[pivotIndex]

	// Partitioning
	l, r := left, right
	for l <= r {
		for nums[l] < pivot {
			l++
		}
		for nums[r] > pivot {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	// Parallel execution only for large subarrays
	if right-left > threshold {
		wg.Add(2)
		go func() {
			defer wg.Done()
			parallelQuickSort(nums, left, r, wg)
		}()
		go func() {
			defer wg.Done()
			parallelQuickSort(nums, l, right, wg)
		}()
	} else {
		// Use regular quicksort for small subarrays
		parallelQuickSort(nums, left, r, wg)
		parallelQuickSort(nums, l, right, wg)
	}
}

func MergeSort(nums []int, left, right int) {
	// One or fewer numbers in the sort range
	if left >= right {
		return
	}

	// Split
	mid := left + (right-left)/2
	MergeSort(nums, left, mid)
	MergeSort(nums, mid+1, right)

	// Merge
	buf := make([]int, right-left+1)
	l, r, p := left, mid+1, 0
	for l <= mid && r <= right {
		if nums[l] <= nums[r] {
			buf[p] = nums[l]
			p, l = p+1, l+1
		} else {
			buf[p] = nums[r]
			p, r = p+1, r+1
		}
	}
	for l <= mid {
		buf[p] = nums[l]
		p, l = p+1, l+1
	}
	for r <= right {
		buf[p] = nums[r]
		p, r = p+1, r+1
	}

	// Write back
	for i, j := left, 0; j < len(buf); i, j = i+1, j+1 {
		nums[i] = buf[j]
	}
}

// Stack entry representing subarray indices
type stackEntry struct {
	left, right int
}

// MergeSortNR performs an iterative (non-recursive) merge sort using a stack.
func MergeSortNR(nums []int) {
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

// merge function merges two sorted subarrays: nums[left: mid+1] and nums[mid+1: right+1]
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

	// Copy merged result back to the original array
	copy(nums[left:right+1], buf)
}

//
// Extension
//

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}

	l, r := left, right
	rg := rand.New(rand.NewSource(time.Now().UnixNano()))
	pivotIndex := rg.Intn(right-left) + left
	pivot := nums[pivotIndex]

	for l <= r {
		for nums[l] < pivot {
			l++
		}
		for nums[r] > pivot {
			r--
		}
		// If pointers crossed, partition is complete
		if l >= r {
			break
		}
		// Swap and move pointers
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

	// left partition [left, r], right partition [r+1, right]
	lsz := r - left + 1
	if lsz >= k {
		return quickSelect(nums, left, r, k)
	} else {
		return quickSelect(nums, r+1, right, k-lsz)
	}
}

func KthSmallestNumber(nums []int, k int) (int, error) {
	if k <= 0 || k > len(nums) {
		return 0, errors.New("k is out of bounds")
	}

	// Make a copy to avoid modifying the original slice
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)

	return quickSelect(numsCopy, 0, len(numsCopy)-1, k), nil
}
