package sort

import (
	"errors"
	"fmt"
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

//
// Extension
//

func TestKthSmallestNumber(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		// Normal cases
		{
			name: "k=1 in sorted array",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    1,
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "k=3 in unsorted array",
			args: args{
				nums: []int{5, 3, 1, 4, 2},
				k:    3,
			},
			want:    3,
			wantErr: nil,
		},
		{
			name: "k=last element",
			args: args{
				nums: []int{9, 2, 7, 1, 4, 6, 8, 3, 5},
				k:    9,
			},
			want:    9,
			wantErr: nil,
		},
		{
			name: "array with duplicates",
			args: args{
				nums: []int{2, 2, 1, 1, 3, 3},
				k:    4,
			},
			want:    2,
			wantErr: nil,
		},
		{
			name: "single element array",
			args: args{
				nums: []int{42},
				k:    1,
			},
			want:    42,
			wantErr: nil,
		},

		// Error cases
		{
			name: "k=0 (invalid)",
			args: args{
				nums: []int{1, 2, 3},
				k:    0,
			},
			want:    0,
			wantErr: errors.New("k is out of bounds"),
		},
		{
			name: "k larger than array length",
			args: args{
				nums: []int{1, 2, 3},
				k:    4,
			},
			want:    0,
			wantErr: errors.New("k is out of bounds"),
		},
		{
			name: "empty array with k=1",
			args: args{
				nums: []int{},
				k:    1,
			},
			want:    0,
			wantErr: errors.New("k is out of bounds"),
		},
		{
			name: "negative k",
			args: args{
				nums: []int{1, 2, 3},
				k:    -1,
			},
			want:    0,
			wantErr: errors.New("k is out of bounds"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the input slice to verify it's not modified
			originalNums := make([]int, len(tt.args.nums))
			copy(originalNums, tt.args.nums)

			got, err := KthSmallestNumber(tt.args.nums, tt.args.k)

			// Check if the input slice was modified
			if len(tt.args.nums) > 0 && len(originalNums) > 0 {
				for i := range tt.args.nums {
					if tt.args.nums[i] != originalNums[i] {
						t.Errorf("input slice was modified at index %d", i)
					}
				}
			}

			if got != tt.want {
				t.Errorf("KthSmallestNumber() got = %v, want %v", got, tt.want)
			}

			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("KthSmallestNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("KthSmallestNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Additional test for random cases
func TestKthSmallestNumberRandom(t *testing.T) {
	rand.Seed(42) // Fixed seed for reproducibility
	for i := 0; i < 100; i++ {
		n := rand.Intn(100) + 1 // Array size between 1 and 100
		nums := make([]int, n)
		for j := range nums {
			nums[j] = rand.Intn(1000)
		}
		k := rand.Intn(n) + 1 // k between 1 and n

		// Test against the naive approach (sorting)
		sorted := make([]int, n)
		copy(sorted, nums)
		sort.Ints(sorted)
		expected := sorted[k-1]

		result, err := KthSmallestNumber(nums, k)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result != expected {
			t.Errorf("for nums=%v, k=%d: expected %d, got %d", nums, k, expected, result)
		}
	}
}

func BenchmarkKthSmallestNumber(b *testing.B) {
	// Generate test data
	sizes := []int{100, 1000, 10000, 100000}
	rand.Seed(time.Now().UnixNano())

	for _, size := range sizes {
		// Create a slice with random numbers
		nums := make([]int, size)
		for i := range nums {
			nums[i] = rand.Intn(size * 10)
		}
		k := size / 2 // Benchmark middle element selection

		// Convert size to string properly
		sizeStr := fmt.Sprintf("%d", size)

		// Reset timer before actual benchmark
		b.Run("Size-"+sizeStr, func(b *testing.B) {
			numsCopy := make([]int, len(nums))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				copy(numsCopy, nums) // Copy inside the loop to ensure fair comparison
				_, _ = KthSmallestNumber(numsCopy, k)
			}
		})

		// Compare with standard library sort for reference
		b.Run("Sort-"+sizeStr, func(b *testing.B) {
			numsCopy := make([]int, len(nums))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				copy(numsCopy, nums)
				sort.Ints(numsCopy)
				_ = numsCopy[k-1]
			}
		})
	}
}
