package leetcode

import "container/list"

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

func preorderTraversal_1(root *TreeNode) (res []int) {

	if root == nil {
		return nil
	}

	var stack = list.New()

	stack.PushBack(root.Right)
	stack.PushBack(root.Left)
	res = append(res, root.Val)

	for stack.Len() > 0 {

		e := stack.Back()
		stack.Remove(e)
		node := e.Value.(*TreeNode) //e是Element类型，其值为e.Value.由于Value为接口，所以要断言

		if node == nil {
			continue
		}

		res = append(res, node.Val)
		stack.PushBack(node.Right)
		stack.PushBack(node.Left)

	}

	return res

}
