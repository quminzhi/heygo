package tree

import (
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
