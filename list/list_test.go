package list

import (
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Reverse middle portion (1->2->3->4->5, left=2, right=4)",
			args: args{
				head:  buildList([]int{1, 2, 3, 4, 5}),
				left:  2,
				right: 4,
			},
			want: buildList([]int{1, 4, 3, 2, 5}),
		},
		{
			name: "Reverse entire list (1->2->3, left=1, right=3)",
			args: args{
				head:  buildList([]int{1, 2, 3}),
				left:  1,
				right: 3,
			},
			want: buildList([]int{3, 2, 1}),
		},
		{
			name: "Reverse single node (5, left=1, right=1)",
			args: args{
				head:  buildList([]int{5}),
				left:  1,
				right: 1,
			},
			want: buildList([]int{5}),
		},
		{
			name: "Reverse first two nodes (1->2->3->4, left=1, right=2)",
			args: args{
				head:  buildList([]int{1, 2, 3, 4}),
				left:  1,
				right: 2,
			},
			want: buildList([]int{2, 1, 3, 4}),
		},
		{
			name: "Reverse last two nodes (1->2->3->4, left=3, right=4)",
			args: args{
				head:  buildList([]int{1, 2, 3, 4}),
				left:  3,
				right: 4,
			},
			want: buildList([]int{1, 2, 4, 3}),
		},
		{
			name: "Empty list (nil, left=1, right=1)",
			args: args{
				head:  nil,
				left:  1,
				right: 1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBetween(tt.args.head, tt.args.left, tt.args.right)
			wantSlice := listToSlice(tt.want)
			if !compareListWithSlice(got, wantSlice) {
				t.Errorf("reverseBetween() = %v, want %v", listToSlice(got), wantSlice)
			}
		})
	}
}

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name     string
		args     args
		expected []int // Expected result as a slice
	}{
		{
			name:     "Empty list",
			args:     args{head: buildList([]int{})},
			expected: []int{},
		},
		{
			name:     "Single node",
			args:     args{head: buildList([]int{1})},
			expected: []int{1},
		},
		{
			name:     "Even-length list",
			args:     args{head: buildList([]int{1, 2, 3, 4})},
			expected: []int{1, 4, 2, 3}, // Reordered: 1 -> 4 -> 2 -> 3
		},
		{
			name:     "Odd-length list",
			args:     args{head: buildList([]int{1, 2, 3, 4, 5})},
			expected: []int{1, 5, 2, 4, 3}, // Reordered: 1 -> 5 -> 2 -> 4 -> 3
		},
		{
			name:     "Two nodes",
			args:     args{head: buildList([]int{1, 2})},
			expected: []int{1, 2}, // No change
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the original list for comparison
			original := listToSlice(tt.args.head)
			// Run the function
			reorderList(tt.args.head)
			// Check if the reordered list matches expected
			got := tt.args.head
			if !compareListWithSlice(got, tt.expected) {
				t.Errorf("reorderList() = %v, expected %v (original: %v)",
					listToSlice(got),
					tt.expected, original)
			}
		})
	}
}

func Test_insertionSortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Empty list",
			args: args{head: buildList([]int{4, 2, 1, 3})},
			want: buildList([]int{1, 2, 3, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Run the function
			got := insertionSortList(tt.args.head)
			if !compareListWithSlice(got, listToSlice(tt.want)) {
				t.Errorf("insertionSortList() = %v, want %v", listToSlice(got),
					listToSlice(tt.want))
			}
		})
	}
}

// Helper function to build a linked list from a slice of integers
func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

// Helper function to compare a linked list with a slice of expected values
func compareListWithSlice(head *ListNode, want []int) bool {
	current := head
	i := 0
	for current != nil {
		if i >= len(want) || current.Val != want[i] {
			return false
		}
		current = current.Next
		i++
	}
	return i == len(want) // Ensure no extra nodes are present
}

// Helper function to convert a linked list to a slice (for error messages)
func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
