package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dumpHead := &ListNode{}
	dumpHead.Next = head
	tmp := dumpHead
	for tmp.Next != nil && tmp.Next.Next != nil {
		tmp.Next = head.Next
		tmp = tmp.Next
		next := tmp.Next
		tmp.Next = head
		tmp = tmp.Next
		tmp.Next = next
		head = next
	}
	return dumpHead.Next

}

// 递归
func swapPairs_(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead

}
