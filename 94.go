package leetcode

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (res []int) {

	var i func(node *TreeNode)

	i = func(node *TreeNode) {

		if node == nil {
			return
		}

		i(node.Left)
		res = append(res, node.Val)
		i(node.Right)

		return
	}

	i(root)

	return
}
