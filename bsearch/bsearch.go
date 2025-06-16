package bsearch

// Binary search can be applied on a sequence with binary nature
//
// l                                  r
// o---------------oo-----------------o
//     attr 1            attr 2
//
// bsearch can be used to find the right end of attr 1 or the left end of
// attr 2.
//
// Update: when l = mid, mid = (l + r + 1) / 2 (assumed that l = r - 1)

// SearchRange finds the first and last positions of a target value in a sorted array.
// The array must be sorted in non-decreasing order for the binary search to work correctly.
// If the target is not found, it returns [-1, -1].
//
// Parameters:
//
//	nums   : A slice of integers sorted in non-decreasing order
//	target : The integer value to search for in the slice
//
// Returns:
//
//	A slice of two integers representing the first and last index of the target value,
//	or [-1, -1] if the target is not found.
func SearchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}

	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l >= 0 && nums[l] != target {
		// target not found
		return result
	}
	result = []int{l, l}

	l, r = 0, len(nums)-1
	for l < r {
		mid := l + (r-l+1)/2
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	result[1] = l

	return result
}

// SearchInRotatedSortedArray finds the index of a target value in a rotated
// sorted array with distinct values.
// The array is sorted in ascending order but may have been rotated at an unknown pivot point.
//
// Parameters:
//
//	nums   : A slice of distinct integers that was sorted in ascending order
//	         and then possibly rotated at an unknown pivot
//	target : The integer value to search for in the slice
//
// Returns:
//
//	The index of the target value if found, or -1 if not found
//
// Approach:
//
//	Uses a modified binary search that first determines which half of the array
//	is properly sorted, then checks whether the target lies within the sorted half.
//	This allows for efficient O(log n) search even with rotation.
func SearchInRotatedSortedArray(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// Find the boundary of two halves
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l+1)/2
		if nums[mid] >= nums[0] {
			l = mid
		} else {
			r = mid - 1
		}
	}
	leftEnd := l

	// Find the target in the right half
	if target >= nums[0] {
		l, r = 0, leftEnd
	} else {
		l, r = leftEnd+1, len(nums)-1
	}
	if l > r {
		// right half doesn't exist
		return -1
	}

	for l < r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l >= 0 && nums[l] == target {
		return l
	} else {
		return -1
	}
}
