package sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func generateRandomSlice(size int, maxVal int) []int {
	rg := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rg.Intn(maxVal)
	}
	return arr
}

func TestQuickSort(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "single element",
			args: args{
				nums:  []int{1},
				left:  0,
				right: 0,
			},
			want: []int{1},
		},
		{
			name: "already sorted",
			args: args{
				nums:  []int{1, 2, 3},
				left:  0,
				right: 2,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "reverse sorted",
			args: args{
				nums:  []int{3, 2, 1},
				left:  0,
				right: 2,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "random order",
			args: args{
				nums:  []int{5, 3, 8, 4, 2},
				left:  0,
				right: 4,
			},
			want: []int{2, 3, 4, 5, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.args.nums))
			copy(input, tt.args.nums)
			QuickSort(input, tt.args.left, tt.args.right)
			if !reflect.DeepEqual(input, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", input, tt.want)
			}
		})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	const size = 1000000
	nums := generateRandomSlice(size, size) // Pre-generate random slice
	b.ResetTimer()                          // Start measuring from this point
	for i := 0; i < b.N; i++ {
		input := make([]int, len(nums))
		copy(input, nums) // Ensure sorting the same data
		QuickSort(input, 0, len(input)-1)
	}
}

// ...........................................

func TestQuickSortNR(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
	}{
		{"Already sorted", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"Reverse sorted", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{"All elements same", []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5}},
		{"Random order", []int{3, 7, 8, 5, 2, 1, 9, 6, 4}},
		{"Empty array", []int{}},
		{"Single element", []int{42}},
		{"Two elements swapped", []int{2, 1}},
		{"Duplicate elements", []int{4, 2, 4, 3, 1, 2, 5, 3, 1, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := append([]int(nil), tc.arr...)
			sort.Ints(expected) // Go’s built-in sorting

			QuickSortNR(tc.arr)

			if len(tc.arr) == 0 && len(expected) == 0 {
				return // Both are empty, no need to check further
			}

			if !reflect.DeepEqual(tc.arr, expected) {
				t.Errorf("Test %s failed: got %v, want %v", tc.name, tc.arr, expected)
			}
		})
	}
}

func BenchmarkQuickSortNR(b *testing.B) {
	const size = 1000000
	nums := generateRandomSlice(size, size) // Pre-generate random slice
	b.ResetTimer()                          // Start measuring from this point
	for i := 0; i < b.N; i++ {
		input := make([]int, len(nums))
		copy(input, nums) // Ensure sorting the same data
		QuickSortNR(input)
	}
}

// ...........................................

func TestParallelQuickSort(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
	}{
		{"Already sorted", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"Reverse sorted", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{"All elements same", []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5}},
		{"Random order", []int{3, 7, 8, 5, 2, 1, 9, 6, 4}},
		{"Empty array", []int{}},
		{"Single element", []int{42}},
		{"Two elements swapped", []int{2, 1}},
		{"Duplicate elements", []int{4, 2, 4, 3, 1, 2, 5, 3, 1, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := append([]int(nil), tc.arr...)
			sort.Ints(expected) // Use Go’s built-in sorting as reference

			ParallelQuickSort(tc.arr)

			if len(tc.arr) == 0 && len(expected) == 0 {
				return // Both are empty, no need to check further
			}

			if !reflect.DeepEqual(tc.arr, expected) {
				t.Errorf("Test %s failed: got %v, want %v", tc.name, tc.arr, expected)
			}
		})
	}
}

func TestParallelQuickSortLargeArray(t *testing.T) {
	const size = 1000000 // Large test case
	arr := generateRandomSlice(size, size)

	expected := append([]int(nil), arr...)
	sort.Ints(expected) // Go’s built-in sorting for correctness check

	ParallelQuickSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("ParallelQuickSort failed on large random array")
	}
}

func BenchmarkParallelQuickSort(b *testing.B) {
	const size = 1000000
	nums := generateRandomSlice(size, size) // Pre-generate random slice
	b.ResetTimer()                          // Start measuring from this point
	for i := 0; i < b.N; i++ {
		input := make([]int, len(nums))
		copy(input, nums) // Ensure sorting the same data
		ParallelQuickSort(input)
	}
}

// ...........................................

func TestMergeSort(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "single element",
			args: args{
				nums:  []int{1},
				left:  0,
				right: 0,
			},
			want: []int{1},
		},
		{
			name: "already sorted",
			args: args{
				nums:  []int{1, 2, 3},
				left:  0,
				right: 2,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "reverse sorted",
			args: args{
				nums:  []int{3, 2, 1},
				left:  0,
				right: 2,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "random order",
			args: args{
				nums:  []int{5, 3, 8, 4, 2},
				left:  0,
				right: 4,
			},
			want: []int{2, 3, 4, 5, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Copy the array to avoid modifying the original data.
			input := make([]int, len(tt.args.nums))
			copy(input, tt.args.nums)
			MergeSort(input, tt.args.left, tt.args.right)
			if !reflect.DeepEqual(input, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", input, tt.want)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	size := 1000000
	nums := generateRandomSlice(size, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := make([]int, len(nums))
		copy(input, nums)
		MergeSort(input, 0, len(input)-1)
	}
}

// ...........................................

func TestMergeSortNR(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
	}{
		{"Already sorted", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"Reverse sorted", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{"All elements same", []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5}},
		{"Random order", []int{3, 7, 8, 5, 2, 1, 9, 6, 4}},
		{"Empty array", []int{}},
		{"Single element", []int{42}},
		{"Two elements swapped", []int{2, 1}},
		{"Duplicate elements", []int{4, 2, 4, 3, 1, 2, 5, 3, 1, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := append([]int(nil), tc.arr...)
			sort.Ints(expected) // Go’s built-in sorting

			MergeSortNR(tc.arr)

			if len(tc.arr) == 0 && len(expected) == 0 {
				return // Both are empty, no need to check further
			}

			if !reflect.DeepEqual(tc.arr, expected) {
				t.Errorf("Test %s failed: got %v, want %v", tc.name, tc.arr, expected)
			}
		})
	}
}

func BenchmarkMergeSortNR(b *testing.B) {
	size := 1000000
	nums := generateRandomSlice(size, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := make([]int, len(nums))
		copy(input, nums)
		MergeSortNR(input)
	}
}
