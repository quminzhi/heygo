package tree

import (
	"fmt"
	"math"
)

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
	queue := make([]*TreeNode, 0)
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	queue = append(queue, root)
	for len(queue) > 0 {
		subRes := make([]int, 0)
		num := len(queue) // The number of nodes in the same level
		for i := 0; i < num; i++ {
			node := queue[0]
			queue = queue[1:]
			subRes = append(subRes, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
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
	var maxLengthToLeafNodeFrom func(root *TreeNode) int
	maxLengthToLeafNodeFrom = func(root *TreeNode) int {
		if root == nil {
			return -1
		}
		left := maxLengthToLeafNodeFrom(root.Left)
		right := maxLengthToLeafNodeFrom(root.Right)
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

	maxLengthToLeafNodeFrom(root)
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

// 110
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	balanced := true

	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftDepth := maxDepth(root.Left)
		rightDepth := maxDepth(root.Right)
		// Check if it is balanced
		if math.Abs(float64(leftDepth-rightDepth)) > 1 {
			balanced = false
		}

		maxInt := func(x, y int) int {
			if x > y {
				return x
			}
			return y
		}
		return maxInt(leftDepth, rightDepth) + 1
	}

	maxDepth(root)
	return balanced
}

// 617
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)

	return root1
}

// 100
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 112
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// Leaf node
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	left := root.Left != nil && hasPathSum(root.Left, targetSum-root.Val)
	right := root.Right != nil && hasPathSum(root.Right, targetSum-root.Val)
	return left || right
}

// 111
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// Leaf node
	if root.Left == nil && root.Right == nil {
		return 1
	}

	minInt := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	if root.Left != nil && root.Right != nil {
		left := minDepth(root.Left)
		right := minDepth(root.Right)
		return minInt(left, right) + 1
	}

	if root.Left != nil {
		// right is nil
		return minDepth(root.Left) + 1
	} else {
		// left is nil
		return minDepth(root.Right) + 1
	}
}

// 236
// Search from bottom to top,
// the first root of subtree that includes both p and q is the LCA
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var lcaNode *TreeNode = nil

	var search func(root, p, q *TreeNode) (state int)
	// state:
	// 0b00, not p nor q found
	// 0b01, p found
	// 0b10, q found
	// 0b11, p and q found
	search = func(root, p, q *TreeNode) (state int) {
		if root == nil {
			return 0
		}
		state = 0
		// Root is p or q
		if p == root {
			state = state | 0b01
		}
		if q == root {
			state = state | 0b10
		}
		// Left search p and q
		state = state | search(root.Left, p, q)
		// Right search p and q
		state = state | search(root.Right, p, q)

		if lcaNode == nil && state == 0b11 {
			lcaNode = root
		}

		return state
	}

	search(root, p, q)
	return lcaNode
}

// 113
func pathSum(root *TreeNode, targetSum int) [][]int {
	paths := make([][]int, 0)

	var dfs func(root *TreeNode, sum int, path []int)
	dfs = func(root *TreeNode, sum int, path []int) {
		if root == nil {
			return
		}
		sum += root.Val
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			if targetSum == sum {
				// path is reference
				pathCopy := append([]int(nil), path...)
				paths = append(paths, pathCopy)
			}
			return
		}
		if root.Left != nil {
			dfs(root.Left, sum, path)
		}
		if root.Right != nil {
			dfs(root.Right, sum, path)
		}
	}

	path := make([]int, 0)
	dfs(root, 0, path)
	return paths
}

// 437
// 1-d prefix + hash => tree-like prefix + hash
//
// ------j-------i-----
// prefix[i] - prefix[j] = targetSum
// prefix[j] = targetSum - prefix[i]
func pathSumIII(root *TreeNode, targetSum int) int {
	res := 0
	cntSum := make(map[int]int) // [sum] -> occurrence

	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum += root.Val
		res += cntSum[sum-targetSum]
		cntSum[sum]++

		if root.Left != nil {
			dfs(root.Left, sum)
		}
		if root.Right != nil {
			dfs(root.Right, sum)
		}
		cntSum[sum]--
	}

	cntSum[0]++ // Virtual node for prefix
	dfs(root, 0)
	return res
}

// 129
func sumNumbers(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}

	var dfs func(root *TreeNode, num int)
	dfs = func(root *TreeNode, num int) {
		if root == nil {
			return
		}
		num = root.Val + num*10
		if root.Left == nil && root.Right == nil {
			res += num
		}
		if root.Left != nil {
			dfs(root.Left, num)
		}
		if root.Right != nil {
			dfs(root.Right, num)
		}
	}

	dfs(root, 0)
	return res
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 117
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == n-1 {
				node.Next = nil
			} else {
				node.Next = queue[0]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

// 98
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(root *TreeNode) (bool, int, int)
	dfs = func(root *TreeNode) (valid bool, minVal int, maxVal int) {
		if root.Left == nil && root.Right == nil {
			return true, root.Val, root.Val
		}
		minInt := func(x, y int) int {
			if x < y {
				return x
			}
			return y
		}
		maxInt := func(x, y int) int {
			if x > y {
				return x
			}
			return y
		}
		minVal, maxVal = root.Val, root.Val
		if root.Left != nil {
			validLeft, minValLeft, maxValLeft := dfs(root.Left)
			if !validLeft || maxValLeft >= root.Val {
				return false, 0, 0
			}
			minVal = minInt(minVal, minValLeft)
			maxVal = maxInt(maxVal, maxValLeft)
		}
		if root.Right != nil {
			validRight, minValRight, maxValRight := dfs(root.Right)
			if !validRight || minValRight <= root.Val {
				return false, 0, 0
			}
			minVal = minInt(minVal, minValRight)
			maxVal = maxInt(maxVal, maxValRight)
		}
		return true, minVal, maxVal
	}

	valid, _, _ := dfs(root)
	return valid
}
