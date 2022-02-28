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

// 递归
func reverseList_(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
