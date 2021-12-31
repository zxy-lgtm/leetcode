package exercise

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	res := &ListNode{0, head}
	tmp := res
	num := 1000
	for tmp.Next != nil && tmp.Next.Next != nil {
		if tmp.Next.Val == tmp.Next.Next.Val {
			num = tmp.Next.Val
			for tmp.Next != nil && tmp.Next.Val == num {
				tmp.Next = tmp.Next.Next
			}
		} else {
			tmp = tmp.Next
		}
	}
	return res.Next
}
