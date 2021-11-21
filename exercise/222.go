package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftNum := countNodes(root.Left)
	rightNum := countNodes(root.Right)
	num := leftNum + rightNum + 1

	return num
}
