package leetcode

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) (dep int) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 {
		l := queue.Len()
		for i := 0; i < l; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			for j := 0; j < len(node.Children); j++ {
				queue.PushBack(node.Children[j])
			}
		}
		dep++
	}

	return
}
