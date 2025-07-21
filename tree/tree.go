package tree

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
