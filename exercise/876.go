package leetcode

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//算长度然后取中间值
func middleNode(head *ListNode) *ListNode {
	var lenth func(*ListNode) int
	lenth = func(head *ListNode) int {
		len := 0
		for head.Next != nil {
			len++
			head = head.Next
		}
		return len
	}
	p := head
	len := lenth(p)
	var end int

	if len%2 == 0 {
		end = len / 2
	} else {
		end = len/2 + 1
	}

	for end > 0 {
		end--
		head = head.Next
	}

	return head

}

//快慢指针
func middleNode2(head *ListNode) *ListNode {
	for front := head; front != nil && front.Next != nil; {
		front, head = front.Next.Next, head.Next
	}
	return head
}