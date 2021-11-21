package leetcode

import "container/list"

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

func inorderTraversal_1(root *TreeNode) (res []int) {

	if root == nil {
		return nil
	}

	stack := list.New()
	node := root
	//先将所有左节点找到，加入栈中
	for node != nil {
		stack.PushBack(node)
		node = node.Left
	}
	//其次对栈中的每个节点先弹出加入到结果集中，再找到该节点的右节点的所有左节点加入栈中
	for stack.Len() > 0 {

		e := stack.Back()
		node := e.Value.(*TreeNode)
		stack.Remove(e)
		//找到该节点的右节点，再搜索他的所有左节点加入栈中
		res = append(res, node.Val)
		node = node.Right

		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}

	}

	return
}
