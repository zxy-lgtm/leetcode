package exercise

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, late := head, head
	for fast.Next != nil && fast.Next.Next != nil && late.Next != nil {
		fast = fast.Next.Next
		late = late.Next
		if fast == late {
			return true
		}
	}
	return false
}
