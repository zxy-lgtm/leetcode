package leetcode

import "container/list"

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

func postorderTraversal_1(root *TreeNode) (res []int) {

	if root == nil {
		return nil
	}

	var stack = list.New()

	stack.PushBack(root.Left)
	stack.PushBack(root.Right)
	res = append(res, root.Val)

	for stack.Len() > 0 {

		e := stack.Back()
		stack.Remove(e)
		node := e.Value.(*TreeNode) //e是Element类型，其值为e.Value.由于Value为接口，所以要断言

		if node == nil {
			continue
		}

		res = append(res, node.Val)
		stack.PushBack(node.Left)
		stack.PushBack(node.Right)

	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	return res
}
