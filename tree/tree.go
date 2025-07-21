package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 145
func postorderTraversal(root *TreeNode) []int {
	post := make([]int, 0)

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		post = append(post, root.Val)
	}

	dfs(root)
	return post
}

// 94
func inorderTraversal(root *TreeNode) []int {
	in := make([]int, 0)

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		in = append(in, root.Val)
		dfs(root.Right)
	}

	dfs(root)
	return in
}

// 144
func preorderTraversal(root *TreeNode) []int {
	pre := make([]int, 0)

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		pre = append(pre, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	return pre
}

// 102
func levelOrder(root *TreeNode) [][]int {
	stk := make([]*TreeNode, 0)
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	stk = append(stk, root)
	for len(stk) > 0 {
		subRes := make([]int, 0)
		num := len(stk) // The number of nodes in the same level
		for i := 0; i < num; i++ {
			node := stk[0]
			stk = stk[1:]
			subRes = append(subRes, node.Val)
			if node.Left != nil {
				stk = append(stk, node.Left)
			}
			if node.Right != nil {
				stk = append(stk, node.Right)
			}
		}
		res = append(res, subRes)
	}
	return res
}

// 108
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

// 105
// preorder: mid, left set, right set
// inorder:  left set, mid, right set
func buildTreeFromPreorderInorder(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	// Guarantee that preorder and inorder consist of unique values
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	// Size of left set == i
	// Left:  preorder[1:i+1] inorder[0:i]
	// Right: preorder[i+1:]  inorder[i+1:]
	root.Left = buildTreeFromPreorderInorder(preorder[1:i+1], inorder[:i])
	root.Right = buildTreeFromPreorderInorder(preorder[i+1:], inorder[i+1:])
	return root
}

// 106
// inorder:   left set, mid, right set
// postorder: left set, right set, mid
func buildTreeFromInorderPostorder(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			break
		}
	}
	// Size of left set == i
	// Left:  inorder[0:i]  postorder[0:i]
	// Right: inorder[i+1:] postorder[i:len(postorder)-1]
	root.Left = buildTreeFromInorderPostorder(inorder[:i], postorder[:i])
	root.Right = buildTreeFromInorderPostorder(inorder[i+1:],
		postorder[i:len(postorder)-1])
	return root
}

// 114
func flatten(root *TreeNode) {
	for root != nil {
		if root.Left != nil {
			// Insert the right chain of left subtree into the root and root.Right
			last := root.Left
			for last.Right != nil {
				last = last.Right
			}
			last.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		// Go to right subtree
		root = root.Right
	}
}

// 889
// preorder:  mid, left set, right set
// postorder: left set, right set, mid
// preorder[mid+1] is the root of left subtree
// The root of left subtree should be the last element in the left set of
// postorder
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	// Cache the position of each value in a postorder sequence
	pos := make(map[int]int)
	for i, val := range postorder {
		pos[val] = i
	}

	var build func(pre []int, post []int, preStart, preEnd, postStart,
		postEnd int) *TreeNode

	build = func(pre []int, post []int, preStart, preEnd, postStart,
		postEnd int) *TreeNode {
		if preStart > preEnd {
			return nil
		}
		root := &TreeNode{Val: preorder[preStart]}
		if preStart == preEnd {
			return root
		}
		// Location of left subtree root in postorder
		k := pos[preorder[preStart+1]]
		leftSize := k - postStart + 1

		// Left:
		// preorder => preStart+1, preStart+1+leftSize-1
		// postorder => postStart, k
		// Right:
		// preorder => preStart+1+leftSize, preEnd
		// postorder => k+1, postEnd-1
		root.Left = build(preorder, postorder, preStart+1, preStart+leftSize,
			postStart, k)
		root.Right = build(preorder, postorder, preStart+1+leftSize, preEnd,
			k+1, postEnd-1)
		return root
	}

	return build(preorder, postorder, 0, len(preorder)-1, 0, len(postorder)-1)
}

// 104
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxInt := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	return maxInt(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 101 *****
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var isMirror func(p, q *TreeNode) bool
	isMirror = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil || p.Val != q.Val {
			return false
		}
		return isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
	}

	return isMirror(root.Left, root.Right)
}

// 226
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var invert func(root *TreeNode) *TreeNode
	invert = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		invert(root.Left)
		invert(root.Right)
		root.Left, root.Right = root.Right, root.Left
		return root
	}

	return invert(root)
}

// 543
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxDiameter := 0
	var maxLengthToLeafNode func(root *TreeNode) int
	maxLengthToLeafNode = func(root *TreeNode) int {
		if root == nil {
			return -1
		}
		left := maxLengthToLeafNode(root.Left)
		right := maxLengthToLeafNode(root.Right)
		// Update maxDiameter
		diameter := left + right + 2
		if diameter > maxDiameter {
			maxDiameter = diameter
		}
		maxInt := func(x, y int) int {
			if x > y {
				return x
			}
			return y
		}
		return maxInt(left, right) + 1
	}

	maxLengthToLeafNode(root)
	return maxDiameter
}

// 257
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	var treePath func(root *TreeNode) []string
	treePath = func(root *TreeNode) []string {
		if root.Left == nil && root.Right == nil {
			return []string{fmt.Sprintf("%v", root.Val)}
		}

		rootPaths := make([]string, 0)
		if root.Left != nil {
			leftPaths := treePath(root.Left)
			// In for range, leftPath is the copy of element in leftPaths
			// Modification to it does not apply to original one
			for _, leftPath := range leftPaths {
				rootPaths = append(rootPaths, fmt.Sprintf("%d->%s", root.Val,
					leftPath))
			}
		}
		if root.Right != nil {
			rightPaths := treePath(root.Right)
			for _, rightPath := range rightPaths {
				rootPaths = append(rootPaths, fmt.Sprintf("%d->%s", root.Val,
					rightPath))
			}
		}
		return rootPaths
	}

	return treePath(root)
}
