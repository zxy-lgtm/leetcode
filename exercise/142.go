package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	late, fast := head, head
	for fast != nil && fast.Next != nil {
		late = late.Next
		fast = fast.Next.Next
		if fast == late {
			fast = head
			for late != fast {
				fast = fast.Next
				late = late.Next
			}
			return late
		}
	}
	return nil
}
