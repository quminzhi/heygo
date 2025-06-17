package bsearch

import (
	"reflect"
	"testing"
)

//
// SearchInsert
//

func TestSearchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// Found cases
		{
			name: "target found in single-element array",
			args: args{nums: []int{5}, target: 5},
			want: 0,
		},
		{
			name: "target found at beginning",
			args: args{nums: []int{1, 3, 5}, target: 1},
			want: 0,
		},
		{
			name: "target found in middle",
			args: args{nums: []int{1, 3, 5}, target: 3},
			want: 1,
		},
		{
			name: "target found at end",
			args: args{nums: []int{1, 3, 5}, target: 5},
			want: 2,
		},

		// Insert cases
		{
			name: "insert at beginning",
			args: args{nums: []int{2, 4, 6}, target: 1},
			want: 0,
		},
		{
			name: "insert in middle",
			args: args{nums: []int{1, 3, 5}, target: 2},
			want: 1,
		},
		{
			name: "insert at end",
			args: args{nums: []int{1, 3, 5}, target: 6},
			want: 3,
		},
		{
			name: "insert in empty array",
			args: args{nums: []int{}, target: 1},
			want: 0,
		},

		// Edge cases
		{
			name: "insert before duplicate values",
			args: args{nums: []int{1, 3, 3, 5}, target: 2},
			want: 1,
		},
		{
			name: "insert at position of first occurrence",
			args: args{nums: []int{1, 3, 3, 5}, target: 3},
			want: 1,
		},
		{
			name: "large array insertion",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				target: 7,
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSearchInsert(b *testing.B) {
	type args struct {
		nums   []int
		target int
	}

	benchmarks := []struct {
		name string
		args args
	}{
		{
			name: "small array",
			args: args{nums: []int{1, 3, 5}, target: 2},
		},
		{
			name: "medium array",
			args: args{nums: makeRange(1, 1000), target: 500},
		},
		{
			name: "large array",
			args: args{nums: makeRange(1, 100000), target: 99999},
		},
		{
			name: "not found large",
			args: args{nums: makeRange(1, 100000), target: 100001},
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SearchInsert(bm.args.nums, bm.args.target)
			}
		})
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

//
// SearchRange
//

func TestSearchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty array",
			args: args{
				nums:   []int{},
				target: 5,
			},
			want: []int{-1, -1},
		},
		{
			name: "single element matching target",
			args: args{
				nums:   []int{5},
				target: 5,
			},
			want: []int{0, 0},
		},
		{
			name: "single element not matching target",
			args: args{
				nums:   []int{3},
				target: 5,
			},
			want: []int{-1, -1},
		},
		{
			name: "multiple elements with single occurrence",
			args: args{
				nums:   []int{1, 3, 5, 7},
				target: 5,
			},
			want: []int{2, 2},
		},
		{
			name: "multiple occurrences of target",
			args: args{
				nums:   []int{1, 2, 2, 2, 3, 4},
				target: 2,
			},
			want: []int{1, 3},
		},
		{
			name: "target at beginning",
			args: args{
				nums:   []int{1, 1, 2, 3, 4},
				target: 1,
			},
			want: []int{0, 1},
		},
		{
			name: "target at end",
			args: args{
				nums:   []int{1, 2, 3, 4, 4},
				target: 4,
			},
			want: []int{3, 4},
		},
		{
			name: "all elements match target",
			args: args{
				nums:   []int{5, 5, 5, 5},
				target: 5,
			},
			want: []int{0, 3},
		},
		{
			name: "target not found in middle",
			args: args{
				nums:   []int{1, 3, 5, 7},
				target: 4,
			},
			want: []int{-1, -1},
		},
		{
			name: "larger array with multiple targets",
			args: args{
				nums:   []int{0, 1, 2, 3, 4, 4, 4, 5, 6, 6, 7},
				target: 4,
			},
			want: []int{4, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSearchRange(b *testing.B) {
	// Test cases for searchRange
	testCases := []struct {
		name   string
		nums   []int
		target int
	}{
		{"small array", []int{1, 2, 2, 3, 4, 4, 4}, 4},
		{"medium array", makeSortedArray(1000, 42), 42},
		{"large array", makeSortedArray(100000, 99999), 99999},
		{"not found", makeSortedArray(10000, -1), 10001},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SearchRange(tc.nums, tc.target)
			}
		})
	}
}

//
// SearchInRotatedSortedArray
//

func Test_searchInRotatedSortedArray(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty array",
			args: args{
				nums:   []int{},
				target: 5,
			},
			want: -1,
		},
		{
			name: "single element - match",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 0,
		},
		{
			name: "single element - no match",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: -1,
		},
		{
			name: "not rotated - target exists",
			args: args{
				nums:   []int{1, 3, 5, 7, 9},
				target: 5,
			},
			want: 2,
		},
		{
			name: "not rotated - target doesn't exist",
			args: args{
				nums:   []int{1, 3, 5, 7, 9},
				target: 4,
			},
			want: -1,
		},
		{
			name: "rotated once - target in right half",
			args: args{
				nums:   []int{6, 7, 1, 2, 3, 4, 5},
				target: 3,
			},
			want: 4,
		},
		{
			name: "rotated once - target in left half",
			args: args{
				nums:   []int{6, 7, 1, 2, 3, 4, 5},
				target: 7,
			},
			want: 1,
		},
		{
			name: "rotated once - target is pivot",
			args: args{
				nums:   []int{6, 7, 1, 2, 3, 4, 5},
				target: 1,
			},
			want: 2,
		},
		{
			name: "rotated multiple times - target found",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			name: "rotated multiple times - target not found",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "target is first element",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 4,
			},
			want: 0,
		},
		{
			name: "target is last element",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 2,
			},
			want: 6,
		},
		{
			name: "large rotated array - target exists",
			args: args{
				nums:   []int{10, 12, 15, 17, 19, 1, 3, 5, 7, 9},
				target: 5,
			},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInRotatedSortedArray(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInRotatedSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSearchInRotatedSortedArray(b *testing.B) {
	// Test cases for rotated array search
	testCases := []struct {
		name   string
		nums   []int
		target int
	}{
		{"small rotated", []int{4, 5, 6, 1, 2, 3}, 2},
		{"medium rotated", makeRotatedArray(1000, 250), 250},
		{"large rotated", makeRotatedArray(100000, 25000), 25000},
		{"not found rotated", makeRotatedArray(10000, 2500), -1},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SearchInRotatedSortedArray(tc.nums, tc.target)
			}
		})
	}
}

//
// Helper functions to generate test data
//

// makeSortedArray creates a sorted array with target repeated multiple times
func makeSortedArray(size, target int) []int {
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i
	}

	// Ensure the target exists multiple times if it's within range
	if target >= 0 && target < size {
		// Place target at three positions
		if target > 0 {
			nums[target-1] = target
		}
		nums[target] = target
		if target < size-1 {
			nums[target+1] = target
		}
	}
	return nums
}

// makeRotatedArray creates a rotated sorted array with pivot at k
func makeRotatedArray(size, pivot int) []int {
	if pivot <= 0 || pivot >= size {
		pivot = size / 4 // Default pivot position
	}

	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i
	}

	// Ensure pivot is within bounds
	if pivot >= size {
		pivot = size - 1
	}

	// Rotate the array safely
	rotated := append(nums[pivot:], nums[:pivot]...)

	// Ensure pivot value exists in the array
	if pivot < len(rotated) {
		rotated[pivot] = pivot
	}

	return rotated
}

func TestSearchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Single element found", args: args{matrix: [][]int{[]int{1}},
			target: 1}, want: true},
		{name: "Single element not found", args: args{matrix: [][]int{[]int{1}},
			target: 0}, want: false},
		{name: "Single element not found 2", args: args{matrix: [][]int{[]int{1}},
			target: 2}, want: false},
		{name: "Single column found", args: args{matrix: [][]int{[]int{1},
			[]int{3}},
			target: 3}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("SearchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
