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

func detectCycle_(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, late := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		late = late.Next
		if fast == late {
			fast = head
			for fast != late {
				fast = fast.Next
				late = late.Next
			}
			return late
		}
	}

	return nil
}
