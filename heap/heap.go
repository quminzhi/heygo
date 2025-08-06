package heap

import (
	"container/heap"
	"math/rand"
	"time"
)

type MaxHeap []int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 1046
func lastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return 0
	}

	h := &MaxHeap{}
	heap.Init(h)

	for _, stone := range stones {
		heap.Push(h, stone)
	}

	absInt := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for h.Len() > 1 {
		one := heap.Pop(h).(int)
		other := heap.Pop(h).(int)
		heap.Push(h, absInt(one-other))
	}

	return heap.Pop(h).(int)
}

// 215
// kth greatest element
// 1. sort O(nlogn)
// 2. heap O(nlogk)
// 3. quick select O(n), n + n/2 + n/4 + ... <= 2n
func findKthLargest(nums []int, k int) int {
	// Initialize random seed once at the beginning
	rg := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Convert kth largest to kth smallest (1-based index)
	k = len(nums) - k

	var quickSelect func(left, right int) int
	quickSelect = func(left, right int) int {
		if left == right {
			return nums[left]
		}

		// Random pivot selection
		pivotIndex := left + rg.Intn(right-left+1)
		pivot := nums[pivotIndex]

		// Move pivot to end for simpler partitioning
		nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]

		// Partition
		// i points to the first element >= pivot during partition
		i := left
		for j := left; j < right; j++ {
			if nums[j] < pivot {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}

		// Move pivot to its final place
		nums[i], nums[right] = nums[right], nums[i]

		if i == k {
			return nums[i]
		} else if i < k {
			return quickSelect(i+1, right)
		} else {
			return quickSelect(left, i-1)
		}
	}

	return quickSelect(0, len(nums)-1)
}
