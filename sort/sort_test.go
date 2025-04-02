package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

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

func generateRandomArray(size int) []int {
	rg := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rg.Intn(1000)
	}
	return arr
}

func BenchmarkQuickSort(b *testing.B) {
	size := 1000
	nums := generateRandomArray(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := make([]int, len(nums))
		copy(input, nums)
		QuickSort(input, 0, len(input)-1)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	size := 1000
	nums := generateRandomArray(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := make([]int, len(nums))
		copy(input, nums)
		MergeSort(input, 0, len(input)-1)
	}
}
