package exercise

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	p := head
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			p.Next = list2
			list2 = list2.Next
			p = p.Next
		} else {
			p.Next = list1
			list1 = list1.Next
			p = p.Next
		}
	}

	if list1 != nil {
		p.Next = list1
	}

	if list2 != nil {
		p.Next = list2
	}

	return head.Next
}

func merge(list []*ListNode, l, r int) *ListNode {
	if l == r {
		return list[l]
	}

	if l > r {
		return nil
	}

	mid := (l + r) >> 1

	return mergeTwoList(merge(list, l, mid), merge(list, mid+1, r))
}
