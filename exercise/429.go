package leetcode

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) (res [][]int) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		length := queue.Len() //记录当前层的数量
		var tmp []int
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			tmp = append(tmp, node.Val)
			for j := 0; j < len(node.Children); j++ {
				queue.PushBack(node.Children[j])
			}
		}
		res = append(res, tmp)
	}
	return
}
