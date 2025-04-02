package sort

import (
	"math/rand"
	"time"
)

// QuickSort sort an array of numbers in [left, right]
func QuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	// Create a new random number generator with a custom source
	rg := rand.New(rand.NewSource(time.Now().UnixNano()))

	l, r, i := left-1, right+1, rg.Intn(right-left+1)+left
	pivot := nums[i]

	for l < r {
		for {
			l++
			// Find the first number greater than the pivot
			if nums[l] >= pivot {
				break
			}
		}
		for {
			r--
			// Find the first number less than the pivot
			if nums[r] <= pivot {
				break
			}
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}

	QuickSort(nums, left, r)
	QuickSort(nums, r+1, right)
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
