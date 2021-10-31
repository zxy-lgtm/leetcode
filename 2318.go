package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	lenA, lenB := 0, 0
	tmpA, tmpB := headA, headB
	for tmpA != nil {
		tmpA = tmpA.Next
		lenA++
	}
	for tmpB != nil {
		tmpB = tmpB.Next
		lenB++
	}
	tmpA, tmpB = headA, headB
	if lenA >= lenB {
		de := lenA - lenB
		for de != 0 {
			tmpA = tmpA.Next
			de--
		}
	} else {
		de := lenB - lenA
		for de != 0 {
			tmpB = tmpB.Next
			de--
		}
	}
	for tmpA != nil && tmpB != nil {
		if tmpA == tmpB {
			return tmpB
		}
		tmpB = tmpB.Next
		tmpA = tmpA.Next
	}
	return nil
}
