package exercise

import "container/list"

/**
 * Definition for a Node.
 **/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		lenth := queue.Len()
		nums := []*Node{}
		for i := 0; i < lenth; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			nums = append(nums, node)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		for i := 0; i < len(nums); i++ {
			if i == len(nums)-1 {
				nums[i].Next = nil
				break
			}
			nums[i].Next = nums[i+1]
		}
	}
	return root
}

// 去除队列
func connect_(root *Node) *Node {
	start := root
	for start != nil {
		var nextStart, last *Node
		handle := func(cur *Node) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = nextStart
	}
	return root
}
