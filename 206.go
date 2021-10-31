package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	res := &ListNode{head.Val, nil}
	for head.Next != nil {
		node := &ListNode{head.Next.Val, res}
		res = node
		head = head.Next
	}

	return res
}
