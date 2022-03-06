package exercise

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

func isValidBST_(root *TreeNode) bool {
	if root == nil {
		return true
	}

	ok1 := isValidBST(root.Left)
	ok2 := isValidBST(root.Right)
	ok3, ok4 := true, true

	if root.Right != nil {
		ok3 = root.Right.Val > root.Val
	}

	if root.Left != nil {
		ok4 = root.Left.Val < root.Val
	}

	return ok1 && ok2 && ok3 && ok4
}
