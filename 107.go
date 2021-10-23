package leetcode

import "container/list"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) (res [][]int) {

	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)
	var tmpArr []int

	for queue.Len() != 0 {
		len := queue.Len()
		for i := 0; i < len; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}

			tmpArr = append(tmpArr, node.Val)
		}
		res = append(res, tmpArr)
		tmpArr = []int{}
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	return

}
