package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}

	LeftH := maxDepth(root.Left) + 1
	RightH := maxDepth(root.Right) + 1
	if abs(LeftH, RightH) > 1 {
		return false
	}
	return true

}

func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(maxDepth(node.Left), maxDepth(node.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	c := a - b
	if c < 0 {
		return -c
	}
	return c
}
