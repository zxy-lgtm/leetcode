package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var travel func(node *TreeNode, tag int) bool

	travel = func(node *TreeNode, tag int) (ok bool) {
		if node.Left == nil && node.Right == nil && tag == 0 {
			return true
		}
		if node.Left != nil {
			if travel(node.Left, tag-node.Left.Val) {
				return true
			}
		}
		if node.Right != nil {
			if travel(node.Right, tag-node.Right.Val) {
				return true
			}
		}

		return false
	}

	return travel(root, targetSum-root.Val)
}
