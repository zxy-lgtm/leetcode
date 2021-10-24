package leetcode

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var compare func(left *TreeNode, right *TreeNode) bool
	compare = func(left *TreeNode, right *TreeNode) bool {
		if left != nil && right == nil {
			return false
		} else if left == nil && right != nil {
			return false
		} else if left == nil && right == nil {
			return true
		} else if left.Val != right.Val {
			return false
		}

		ok1 := compare(left.Left, right.Right)
		ok2 := compare(left.Right, right.Left)

		return ok1 && ok2
	}

	return compare(root.Left, root.Right)
}
