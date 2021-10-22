package leetcode

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) (res []int) {

	var p func(node *TreeNode)

	p = func(node *TreeNode) {

		if node == nil {
			return
		}

		p(node.Left)
		p(node.Right)
		res = append(res, node.Val)

		return
	}

	p(root)

	return
}
