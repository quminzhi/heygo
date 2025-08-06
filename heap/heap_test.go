package heap

import "testing"

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
