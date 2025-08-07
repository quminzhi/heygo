package heap

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"single", args{nums: []int{1}, k: 1}, 1},
		{"general", args{nums: []int{3, 2, 1, 5, 6, 4}, k: 2}, 5},
		{"general with repeat", args{nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k: 4},
			4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "single", args: args{nums: []int{1}, k: 1}, want: []int{1}},
		{name: "general", args: args{nums: []int{1, 1, 1, 2, 2, 3}, k: 2},
			want: []int{1, 2}},
		{name: "general1", args: args{nums: []int{3, 0, 1, 0}, k: 1},
			want: []int{0}},
		{name: "general2", args: args{nums: []int{5, 3, 1, 1, 1, 3, 73, 1},
			k: 2}, want: []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.args.nums, tt.args.k)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
