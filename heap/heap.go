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

// 347
// 1. heap O(nlogk)
//
//	type Pair struct {
//		num, freq int
//	}
//
// type MinHeap []Pair
//
// func (h MinHeap) Len() int { return len(h) }
//
//	func (h MinHeap) Less(i, j int) bool {
//		return h[i].freq < h[j].freq
//	}
//
//	func (h MinHeap) Swap(i, j int) {
//		h[i], h[j] = h[j], h[i]
//	}
//
//	func (h *MinHeap) Push(x interface{}) {
//		*h = append(*h, x.(Pair))
//	}
//
//	func (h *MinHeap) Pop() interface{} {
//		old := *h
//		n := len(old)
//		x := old[n-1]
//		*h = old[:n-1]
//		return x
//	}
//
//	func topKFrequent(nums []int, k int) []int {
//		m := make(map[int]int)
//		for _, num := range nums {
//			m[num]++
//		}
//
//		h := &MinHeap{}
//		heap.Init(h)
//
//		for num, freq := range m {
//			heap.Push(h, Pair{num, freq})
//			for h.Len() > k {
//				heap.Pop(h)
//			}
//		}
//
//		res := make([]int, 0)
//		for h.Len() > 0 {
//			res = append(res, heap.Pop(h).(Pair).num)
//		}
//
//		return res
//	}
//
// 2. quick select O(n)
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	type Pair struct {
		num, freq int
	}
	f := make([]Pair, 0)
	for num, freq := range m {
		f = append(f, Pair{num, freq})
	}

	// Convert kth largest number (1-based) to its increasing order index (
	// 0-based)
	kth := len(f) - k
	// Quick sort on array f
	rg := rand.New(rand.NewSource(time.Now().UnixNano()))

	var quickSort func(left, right int)
	quickSort = func(left, right int) {
		if left >= right {
			return
		}

		pivotIndex := left + rg.Intn(right-left+1)
		pivot := f[pivotIndex]
		f[right], f[pivotIndex] = f[pivotIndex], f[right]
		i := left
		for j := left; j < right; j++ {
			if f[j].freq < pivot.freq {
				f[i], f[j] = f[j], f[i]
				i++
			}
		}
		f[i], f[right] = f[right], f[i]

		if i == kth {
			// kth largest number found, meaning top k found
			return
		} else if i < kth {
			quickSort(i+1, right)
		} else {
			quickSort(left, i-1)
		}
	}
	quickSort(0, len(f)-1)

	res := make([]int, 0)
	for i, j := 0, len(f)-1; i < k; i++ {
		res = append(res, f[j].num)
		j--
	}

	return res
}

// 692
type Pair struct {
	word string
	freq int
}

type MinHeap []Pair

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	if h[i].freq == h[j].freq {
		return h[i].word > h[j].word
	}
	return h[i].freq < h[j].freq
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequentWords(words []string, k int) []string {
	dict := make(map[string]int)
	for _, word := range words {
		dict[word] = dict[word] + 1
	}
	h := make(MinHeap, 0)
	heap.Init(&h)
	for word, freq := range dict {
		heap.Push(&h, Pair{word, freq})
		for h.Len() > k {
			heap.Pop(&h)
		}
	}

	res := make([]string, 0)
	for h.Len() > 0 {
		res = append(res, heap.Pop(&h).(Pair).word)
	}

	reverse := func(s []string) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}
	reverse(res)

	return res
}
