package exercise

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}

func getIntersectionNode_(headA, headB *ListNode) *ListNode {
	mp := make(map[*ListNode]int)
	p, q := headA, headB
	for p != nil || q != nil {
		if p != nil {
			mp[p]++
			if mp[p] > 1 {
				return p
			}
			p = p.Next
		}

		if q != nil {
			mp[q]++
			if mp[q] > 1 {
				return q
			}
			q = q.Next
		}
	}

	return nil
}

// 双指针
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	fp, fq := p, q
	for fp != nil || fq != nil {
		if fp != nil && fq != nil {
			fp = fp.Next
			fq = fq.Next
		} else if fq == nil {
			fp = fp.Next
			p = p.Next
		} else if fp == nil {
			fq = fq.Next
			q = q.Next
		}
	}

	for p != q && p != nil && q != nil {
		p = p.Next
		q = q.Next
	}

	return p
}
