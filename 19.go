package leetcode

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast, slow := head, head

	for n > 0 {
		fast = fast.Next
		n--
	}
	if fast == nil {
		return head.Next
	}

	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head

}
