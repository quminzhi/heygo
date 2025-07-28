package tree

import (
	"math"
	"reflect"
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "nil root",
			args: args{root: nil},
			want: []string{},
		},
		{
			name: "single node",
			args: args{root: &TreeNode{Val: 1}},
			want: []string{"1"},
		},
		{
			name: "simple tree",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			}},
			want: []string{"1->2->5", "1->3"},
		},
		{
			name: "unbalanced tree",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 6,
						},
					},
				},
			}},
			want: []string{"1->2->4", "1->3->5->6"},
		},
		{
			name: "only left children",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			}},
			want: []string{"1->2->3"},
		},
		{
			name: "only right children",
			args: args{root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			}},
			want: []string{"1->2->3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty tree",
			args: args{root: nil},
			want: true,
		},
		{
			name: "single node",
			args: args{root: &TreeNode{Val: 1}},
			want: true,
		},
		{
			name: "balanced tree with two levels",
			args: args{root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			}},
			want: true,
		},
		{
			name: "unbalanced tree left heavy",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			}},
			want: false,
		},
		{
			name: "unbalanced tree right heavy",
			args: args{root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			}},
			want: false,
		},
		{
			name: "balanced tree with three levels",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			}},
			want: true,
		},
		{
			name: "balanced tree with differing depths but within 1",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 3},
			}},
			want: true,
		},
		{
			name: "unbalanced tree with difference more than 1",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 4},
					},
				},
				Right: &TreeNode{Val: 5},
			}},
			want: false,
		},
		{
			name: "complex balanced tree",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  4,
						Left: &TreeNode{Val: 7},
					},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val:   6,
						Right: &TreeNode{Val: 8},
					},
				},
			}},
			want: false,
		},
		{
			name: "complex unbalanced tree",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  4,
						Left: &TreeNode{Val: 7},
					},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 6,
						Right: &TreeNode{
							Val:   8,
							Right: &TreeNode{Val: 9},
						},
					},
				},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "empty tree",
			args: args{
				root:      nil,
				targetSum: 0,
			},
			want: [][]int{},
		},
		{
			name: "single node matching target",
			args: args{
				root:      &TreeNode{Val: 5},
				targetSum: 5,
			},
			want: [][]int{{5}},
		},
		{
			name: "single node not matching target",
			args: args{
				root:      &TreeNode{Val: 5},
				targetSum: 7,
			},
			want: [][]int{},
		},
		{
			name: "multiple paths with one matching",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   11,
							Left:  &TreeNode{Val: 7},
							Right: &TreeNode{Val: 2},
						},
					},
					Right: &TreeNode{
						Val:  8,
						Left: &TreeNode{Val: 13},
						Right: &TreeNode{
							Val:   4,
							Left:  &TreeNode{Val: 5},
							Right: &TreeNode{Val: 1},
						},
					},
				},
				targetSum: 22,
			},
			want: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			name: "no paths matching target",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
				targetSum: 5,
			},
			want: [][]int{},
		},
		{
			name: "negative numbers in path",
			args: args{
				root: &TreeNode{
					Val: -2,
					Right: &TreeNode{
						Val:   -3,
						Right: &TreeNode{Val: -1},
					},
				},
				targetSum: -6,
			},
			want: [][]int{{-2, -3, -1}},
		},
		{
			name: "multiple paths with multiple matching",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 5},
					},
					Right: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 7},
					},
				},
				targetSum: 7,
			},
			want: [][]int{{1, 2, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty tree",
			args: args{root: nil},
			want: true,
		},
		{
			name: "single node",
			args: args{root: &TreeNode{Val: 1}},
			want: true,
		},
		{
			name: "valid small BST",
			args: args{root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			}},
			want: true,
		},
		{
			name: "invalid small BST - left equals root",
			args: args{root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			}},
			want: false,
		},
		{
			name: "invalid small BST - right equals root",
			args: args{root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			}},
			want: false,
		},
		{
			name: "invalid BST - left greater than root",
			args: args{root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 1},
			}},
			want: false,
		},
		{
			name: "valid complex BST",
			args: args{root: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{
						Val:   6,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 7},
					},
				},
				Right: &TreeNode{
					Val: 10,
					Right: &TreeNode{
						Val:  14,
						Left: &TreeNode{Val: 13},
					},
				},
			}},
			want: true,
		},
		{
			name: "invalid complex BST - violates BST property in subtree",
			args: args{root: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{
						Val:   6,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 9}, // 9 > root 8, should be false
					},
				},
				Right: &TreeNode{
					Val: 10,
					Right: &TreeNode{
						Val:  14,
						Left: &TreeNode{Val: 13},
					},
				},
			}},
			want: false,
		},
		{
			name: "invalid BST with duplicate values",
			args: args{root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 1},
				},
				Right: &TreeNode{Val: 3},
			}},
			want: false,
		},
		{
			name: "valid BST with min and max values",
			args: args{root: &TreeNode{
				Val: math.MaxInt64,
				Left: &TreeNode{
					Val: math.MinInt64,
				},
			}},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInorderTraversalNonRecursive(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty tree",
			args: args{root: nil},
			want: []int{},
		},
		{
			name: "single node",
			args: args{root: &TreeNode{Val: 1}},
			want: []int{1},
		},
		{
			name: "left subtree only",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			}},
			want: []int{3, 2, 1},
		},
		{
			name: "right subtree only",
			args: args{root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			}},
			want: []int{1, 2, 3},
		},
		{
			name: "balanced tree",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			}},
			want: []int{4, 2, 5, 1, 6, 3, 7},
		},
		{
			name: "unbalanced tree",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
					},
				},
			}},
			want: []int{4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversalNonRecursive(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversalNonRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
