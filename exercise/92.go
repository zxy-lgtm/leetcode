package exercise

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseHead(head, right)
	}

	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

var back = &ListNode{}

func reverseHead(head *ListNode, n int) *ListNode {
	if n == 1 {
		back = head.Next
		return head
	}

	last := reverseHead(head.Next, n-1)
	head.Next.Next = head
	head.Next = back
	return last
}
