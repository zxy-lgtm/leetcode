package leetcode

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (res []int) {

	var r func(node *TreeNode)

	r = func(node *TreeNode) {

		if node == nil {
			return
		}

		res = append(res, node.Val)
		r(node.Left)
		r(node.Right)
	}

	r(root)

	return
}
