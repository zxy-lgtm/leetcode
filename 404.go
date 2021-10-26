package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) (total int) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		total += root.Left.Val
	}
	total += sumOfLeftLeaves(root.Left)
	total += sumOfLeftLeaves(root.Right)

	return
}
