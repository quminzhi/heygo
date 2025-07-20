package list

import (
	"reflect"
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

func Test_oddEvenList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty list",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "two elements",
			input:    []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "odd number of elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 3, 5, 2, 4},
		},
		{
			name:     "even number of elements",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: []int{1, 3, 5, 2, 4, 6},
		},
		{
			name:     "longer list",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10},
		},
		{
			name:     "all odd values",
			input:    []int{1, 3, 5, 7, 9},
			expected: []int{1, 5, 9, 3, 7},
		},
		{
			name:     "all even values",
			input:    []int{2, 4, 6, 8, 10},
			expected: []int{2, 6, 10, 4, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputList := buildList(tt.input)
			resultList := oddEvenList(inputList)
			result := listToSlice(resultList)

			if len(result) != len(tt.expected) {
				t.Errorf("expected length %d, got %d", len(tt.expected), len(result))
				return
			}

			for i := range tt.expected {
				if result[i] != tt.expected[i] {
					t.Errorf("at index %d: expected %d, got %d", i, tt.expected[i], result[i])
				}
			}
		})
	}
}

func Test_splitListToParts(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		k        int
		expected [][]int
	}{
		{
			name:     "empty list with k=3",
			input:    []int{},
			k:        3,
			expected: [][]int{{}, {}, {}},
		},
		{
			name:     "single element with k=1",
			input:    []int{1},
			k:        1,
			expected: [][]int{{1}},
		},
		{
			name:     "single element with k=5",
			input:    []int{1},
			k:        5,
			expected: [][]int{{1}, {}, {}, {}, {}},
		},
		{
			name:     "list length equals k",
			input:    []int{1, 2, 3},
			k:        3,
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "list length less than k",
			input:    []int{1, 2, 3},
			k:        5,
			expected: [][]int{{1}, {2}, {3}, {}, {}},
		},
		{
			name:     "list length greater than k with equal parts",
			input:    []int{1, 2, 3, 4, 5, 6},
			k:        3,
			expected: [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:     "list length greater than k with unequal parts",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			k:        3,
			expected: [][]int{{1, 2, 3, 4}, {5, 6, 7}, {8, 9, 10}},
		},
		{
			name:     "large k with small list",
			input:    []int{1, 2, 3},
			k:        10,
			expected: [][]int{{1}, {2}, {3}, {}, {}, {}, {}, {}, {}, {}},
		},
		{
			name:     "k=1 returns original list",
			input:    []int{1, 2, 3, 4, 5},
			k:        1,
			expected: [][]int{{1, 2, 3, 4, 5}},
		},
		{
			name:     "exact division with larger list",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			k:        4,
			expected: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := buildList(tt.input)
			parts := splitListToParts(list, tt.k)
			result := partsToSlices(parts)

			if len(result) != len(tt.expected) {
				t.Errorf("expected %d parts, got %d", len(tt.expected), len(result))
				return
			}

			for i := range tt.expected {
				if !reflect.DeepEqual(result[i], tt.expected[i]) {
					t.Errorf("part %d: expected %v, got %v", i, tt.expected[i], result[i])
				}
			}
		})
	}
}

func Test_MergeKLists(t *testing.T) {
	tests := []struct {
		name     string
		lists    []*ListNode
		expected []int
	}{
		{
			name:     "Empty input",
			lists:    []*ListNode{},
			expected: []int{},
		},
		{
			name:     "Single empty list",
			lists:    []*ListNode{nil},
			expected: []int{},
		},
		{
			name: "Single non-empty list",
			lists: []*ListNode{
				sliceToList([]int{1, 3, 5}),
			},
			expected: []int{1, 3, 5},
		},
		{
			name: "Multiple sorted lists",
			lists: []*ListNode{
				sliceToList([]int{1, 4, 7}),
				sliceToList([]int{2, 5, 8}),
				sliceToList([]int{3, 6, 9}),
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "Lists with varying lengths",
			lists: []*ListNode{
				sliceToList([]int{1, 5}),
				sliceToList([]int{2}),
				sliceToList([]int{3, 4, 6}),
			},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "All lists empty",
			lists: []*ListNode{
				nil,
				nil,
				nil,
			},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeKLists(tt.lists)
			got := listToSlice(result)
			if !slicesEqual(got, tt.expected) {
				t.Errorf("MergeKLists() = %v, want %v", got, tt.expected)
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

// Helper function to convert list of linked lists to slices
func partsToSlices(parts []*ListNode) [][]int {
	var result [][]int
	for _, part := range parts {
		slice := make([]int, 0)
		current := part
		for current != nil {
			slice = append(slice, current.Val)
			current = current.Next
		}
		result = append(result, slice)
	}
	return result
}

// Helper: Converts a slice to a linked list.
func sliceToList(nums []int) *ListNode {
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

// Helper: Checks if two slices are equal.
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
