package exercise

type ListNode struct {
	Val  int
	Next *ListNode
}

//mine
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{0, nil}
	temp := result
	carry := 0
	for l1 != nil || l2 != nil {

		if l1 == nil {
			l1 = &ListNode{0, nil}
		}
		if l2 == nil {
			l2 = &ListNode{0, nil}
		}

		sum := l1.Val + l2.Val + carry
		temp.Val = sum % 10

		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		l1 = l1.Next
		l2 = l2.Next

		if l1 == nil && l2 == nil {
			if carry == 1 {
				temp.Next = &ListNode{0, nil}
				temp = temp.Next
				temp.Val = carry
			}

			break
		}

		temp.Next = &ListNode{0, nil}
		temp = temp.Next

	}
	return result
}

func addTwoNumbers_(l1 *ListNode, l2 *ListNode) *ListNode {
	tag := 0
	head := &ListNode{0, nil}
	res := head
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + tag
		l := &ListNode{sum % 10, nil}
		res.Next = l
		res = l
		tag = sum / 10
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		sum := l1.Val + tag
		l := &ListNode{Val: sum % 10}
		res.Next = l
		res = l
		tag = sum / 10
		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val + tag
		l := &ListNode{Val: sum % 10}
		res.Next = l
		res = l
		tag = sum / 10
		l2 = l2.Next
	}

	if tag != 0 {
		l := &ListNode{Val: tag}
		res.Next = l
	}

	return head.Next
}
