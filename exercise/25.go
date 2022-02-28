package exercise

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newhead := reverseHelp(a, b)
	a.Next = reverseKGroup(b, k)
	return newhead
}

func reverseHelp(head *ListNode, end *ListNode) *ListNode {
	pre := &ListNode{}
	cur, nxt := head, head
	for cur != end {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre

}
