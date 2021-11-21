package leetcode

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) (res int) {
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 {
		l := queue.Len()
		for i := 0; i < l; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if i == 0 {
				res = node.Val
			}
		}
	}
	return
}
