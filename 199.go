package leetcode

import "container/list"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	var res [][]int

	for queue.Len() != 0 {
		len := queue.Len()
		tmp := make([]int, 0)
		for i := 0; i < len; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		res = append(res, tmp)
	}

	for i := 0; i < len(res); i++ {
		ans = append(ans, res[i][len(res[i])-1])
	}
	return
}
