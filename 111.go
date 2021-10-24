package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}

	if root.Right == nil && root.Left != nil {
		return 1 + minDepth(root.Left)
	}

	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}

	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
